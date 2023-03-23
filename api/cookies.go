package api

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	db "Microservice-Login/database/sqlc"
	"Microservice-Login/util"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type GetCookieRequest struct {
	UserName string `json:"user_name" binding:"required"`
	CookieID string `json:"cookie_id" binding:"required"`
}

func (s *Server) SetTokenCookie(ctx *gin.Context, token string) {

	cookie := http.Cookie{
		Name:  "token",
		Value: token,
	}
	http.SetCookie(ctx.Writer, &cookie)
	log.Println("Token Cookie set successfully:", token)
}

func (s *Server) SetUserCookie(ctx *gin.Context, userDetail db.UserDetail, duration time.Duration) {

	expires := time.Now().Add(duration)
	cookieID := getUUIDInt32()
	cookie := http.Cookie{
		Name:    userDetail.UserName,
		Value:   strconv.Itoa(int(cookieID)),
		Expires: expires,
	}
	http.SetCookie(ctx.Writer, &cookie)

	arg := db.InsertCookieParams{
		UserName:  userDetail.UserName,
		CookieID:  cookieID,
		ExpiresAt: expires,
	}

	userCookie, err := s.store.InsertCookie(context.Background(), arg)
	if err != nil {
		fmt.Println("Cannot set cookie")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot set cookie"})
		return
	}
	log.Println("User Cookie set:", userCookie)
}

func (s *Server) GetCookie(r *http.Request, name string) (*http.Cookie, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func (s *Server) ClearTokenCookie(w http.ResponseWriter) {
	expires := time.Unix(0, 0)
	cookie := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expires,
	}
	http.SetCookie(w, &cookie)
}

func (s *Server) ClearUserCookie(w http.ResponseWriter, name string) {
	expires := time.Unix(0, 0)
	cookie := http.Cookie{
		Name:    name,
		Value:   "",
		Expires: expires,
	}

	s.store.DeleteCookieByUserName(context.Background(), name)
	http.SetCookie(w, &cookie)
}

func getUUIDInt32() int32 {
	cookieValue := uuid.NewV4()

	// Convert UUID to int32
	idBytes := cookieValue.Bytes()
	idInt := big.NewInt(0)
	idInt.SetBytes(idBytes)
	idInt32 := int32(idInt.Uint64())

	return idInt32
}

// This middleware checks for the cookie in local storage and in the database so that it ensures the user is logged in.
func (s *Server) AuthCookieMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var cookieReq GetCookieRequest

		if err := ctx.ShouldBindJSON(&cookieReq); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			ctx.Abort()
			return
		}

		cookie, err := s.GetCookie(ctx.Request, cookieReq.UserName)
		if err != nil || cookie == nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			ctx.Abort()
			return
		}

		cookieParts := strings.Split(cookie.Value, ":")
		if len(cookieParts) != 2 {
			fmt.Println("Invalid Cookie for:" + cookieReq.UserName)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid cookie value"})
			ctx.Abort()
			return
		}

		username := cookieParts[0]
		cookieID := util.ConvertToInt32(cookieParts[1])

		dbCookie, err := s.store.SelectCookieByUserName(ctx, username)

		if dbCookie.CookieID != cookieID {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			ctx.Abort()
			return
		}

		ctx.Next()
		//Need to add redirect logic
	}
}

func (s *Server) AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader("authorization")

		if len(authorizationHeader) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is not provided"})
			ctx.Abort()
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {

			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			ctx.Abort()
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unsupported authorization type " + authorizationType})
			ctx.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := s.jwtVerfier.GetMetaData(accessToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		// add userID to the context of the request
		ctx.Set("userID", payload.UserID)
		ctx.Next()
	}
}

// Used for testing if cookies exist if not redirect.
func (server *Server) TestCookie(ctx *gin.Context) {
	var userReq UserLogoutRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// Try to get the cookie from local storage
	cookie, err := server.GetCookie(ctx.Request, userReq.UserName)
	if err != nil || cookie == nil {
		// If the cookie is not found in local storage, try to get it from the database
		dbCookie, err := server.store.SelectCookieByUserName(ctx, userReq.UserName)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		if dbCookie == (db.UserCookie{}) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
			return
		}

		// If the cookie is found in the database, set it in local storage
		// cookie = &http.Cookie{
		// 	Name:   "session",
		// 	Value:  fmt.Sprintf("%s:%d", dbCookie.UserName, dbCookie.CookieID),
		// 	Path:   "/",
		// 	Secure: true,
		// }
		// http.SetCookie(ctx.Writer, cookie)
	}

	ctx.JSON(http.StatusOK, "Testing middleware for "+userReq.UserName)
}

//JED --- NOW need to check how to authenticate the cookies on each render onto a newpage using the middleware. iF NOT redirect to home page
// On logout change the user logout to smth else.

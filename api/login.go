package api

import (
	db "Microservice-Login/database/sqlc"
	"Microservice-Login/util"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type UserLoginRequest struct {
	UserName     string `json:"user_name" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}

type UserUserNameRequest struct {
	UserName string `json:"user_name" binding:"required"`
}

type UserLoginResponse struct {
	Message  string        `json:"message"`
	UserName string        `json:"user_name"`
	Status   int           `json:"response_status"`
	Cookie   http.Cookie   `json:"cookie"`
	Details  db.UserDetail `json:"user_details"`
}

func (server *Server) userLogin(ctx *gin.Context) {

	var userReq UserLoginRequest

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Message: "Missing Input",
		})
		return
	}

	//To ensure username/password has >8 characters (returns bool)
	validatedUserName := util.ValidateUsername(userReq.UserName)
	validatedPassword := util.ValidatePassword(userReq.UserPassword)

	if !validatedUserName {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Message: "Invalid Username Format",
		})
		return
	}

	if !validatedPassword {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Message: "Invalid Password Format",
		})
		return
	}

	//Retrives the credential if does not exist, returns error
	userCred, err := server.store.GetUserByUsername(ctx, userReq.UserName)

	if err != nil {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Message: "Username does not exist",
		})
		fmt.Println(userCred)
		return
	}

	correctPassword := util.CheckStringHash(userReq.UserPassword, userCred.UserPassword)

	if !correctPassword {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Message: "Wrong Password",
		})
		return
	}

	cookieVal := getUUIDInt32cookie()

	// Set cookie with JWT token (passed to front end to set)
	cookie := http.Cookie{
		Name:     userCred.UserName,
		Value:    strconv.Itoa(int(cookieVal)),
		HttpOnly: true,
		Path:     "/",
		Domain:   "localhost",
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(ctx.Writer, &cookie)

	// Set cookies in database
	arg := db.InsertCookieParams{
		UserName: userCred.UserName,
		CookieID: cookieVal,
	}

	_, _ = server.store.InsertCookie(ctx, arg)
	if err != nil {
		fmt.Println("Cannot set cookie")
		ctx.JSON(http.StatusOK, gin.H{"error": "Cannot set cookie"})
		return
	}

	fmt.Println("Set cookies for:" + userCred.UserName)
	ctx.JSON(http.StatusOK, UserLoginResponse{
		Message:  "User logged in",
		UserName: userCred.UserName,
		Status:   http.StatusOK,
		Cookie:   cookie,
	})
}

func (server *Server) userLogout(ctx *gin.Context) {
	var userReq UserUserNameRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// server.ClearCookie(ctx.Writer, userReq.UserName)

	ctx.JSON(http.StatusOK, "Cookies deleted for:"+userReq.UserName)
}

func (server *Server) getUserDetail(ctx *gin.Context) {
	var userReq UserUserNameRequest
	err := ctx.BindJSON(&userReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userCred, err := server.store.GetUserByUsername(ctx, userReq.UserName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, UserLoginResponse{
		Message:  "User data retrieved from backend",
		UserName: userCred.UserName,
		Status:   http.StatusOK,
		Details:  userCred,
	})
}

func getUUIDInt32cookie() int32 {
	cookieValue := uuid.NewV4()

	// Convert UUID to int32
	idBytes := cookieValue.Bytes()
	idInt := big.NewInt(0)
	idInt.SetBytes(idBytes)
	idInt32 := int32(idInt.Uint64())

	return idInt32
}

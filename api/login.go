package api

import (
	"Microservice-Login/util"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	UserName     string `json:"user_name" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}

type userLoginResponse struct {
	Error   error       `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
type UserLogoutRequest struct {
	UserName string `json:"user_name" binding:"required"`
}

func (server *Server) welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Welcome to login service")
}

func (server *Server) TestAuthentication(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "see this page when authorized token is presented in header.")
}

func (server *Server) userLogin(ctx *gin.Context) {

	var userReq UserLoginRequest
	var userRes userLoginResponse

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		userRes = userLoginResponse{
			Error: err,
		}
		ctx.JSON(http.StatusBadRequest, userRes)
		return
	}

	_, err := ctx.Request.Cookie(userReq.UserName)

	if err == nil {
		// Cookie already exists, do not set again
		userRes = userLoginResponse{
			Message: fmt.Sprint("User already logged in:" + userReq.UserName),
		}
		ctx.JSON(http.StatusOK, userRes)
		return
	}

	//To ensure username/password has >8 characters (returns bool)
	validatedUserName := util.ValidateUsername(userReq.UserName)
	validatedPassword := util.ValidatePassword(userReq.UserPassword)

	if !validatedUserName {
		userRes = userLoginResponse{
			Error: errors.New("invalid username format"),
		}
		ctx.JSON(http.StatusBadRequest, userRes)
		return
	}

	if !validatedPassword {
		userRes = userLoginResponse{
			Error: errors.New("invalid password format"),
		}
		ctx.JSON(http.StatusBadRequest, userRes)
		return
	}

	//Retrives the credential if does not exist, returns error
	userCred, err := server.store.GetUserByUsername(ctx, userReq.UserName)
	fmt.Println(userReq.UserName)
	if err != nil {
		userRes = userLoginResponse{
			Error: errors.New("could not get cookie information from server"),
		}
		ctx.JSON(http.StatusBadRequest, userRes)
		return
	}

	correctPassword := util.CheckStringHash(userReq.UserPassword, userCred.UserPassword)

	if !correctPassword {
		userRes = userLoginResponse{
			Error: errors.New("wrong password"),
		}
		ctx.JSON(http.StatusBadRequest, userRes)
		return
	}

	server.SetUserCookie(ctx, userCred, time.Hour)
	log.Println("Set cookies for:" + userCred.UserName)

	// When user hit login entrypoint successfully, assign a new JWT token and set it to cookie
	// The browser needs to carry this cookie in every request header later on.
	duration := time.Duration(time.Second) * time.Duration(server.AppCon.TokenExpireSecs)
	token, err := server.jwtMaker.CreateJWTToken(string(userCred.ID), duration)
	if err != nil {
		log.Fatal(err)
	}
	server.SetTokenCookie(ctx, token)

	data := map[string]string{
		"user":  userCred.UserName,
		"token": token}

	userRes = userLoginResponse{
		Message: "user login successfully",
		Data:    data,
	}

	// Send back to browser so that browser could carry this token in future queries
	ctx.JSON(http.StatusOK, userRes)
}

func (server *Server) userLogout(ctx *gin.Context) {
	var userReq UserLogoutRequest
	var userRes userLoginResponse

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		userRes = userLoginResponse{
			Error: err,
		}
		ctx.JSON(http.StatusBadRequest, userRes)
		return
	}

	_, err := ctx.Request.Cookie(userReq.UserName)

	if err != nil {
		userRes = userLoginResponse{
			Message: fmt.Sprint("User already logged out:" + userReq.UserName),
		}
		ctx.JSON(http.StatusOK, userRes)
		return
	}

	server.ClearUserCookie(ctx.Writer, userReq.UserName)
	server.ClearTokenCookie(ctx.Writer)

	userRes = userLoginResponse{
		Message: fmt.Sprint("login out successfully. Cookies deleted for ", userReq.UserName),
	}
	ctx.JSON(http.StatusOK, userRes)
}

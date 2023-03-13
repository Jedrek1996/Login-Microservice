package api

import (
	"Microservice-Login/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	UserName     string `json:"user_name" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}

type UserLogoutRequest struct {
	UserName string `json:"user_name" binding:"required"`
}

func (server *Server) userLogin(ctx *gin.Context) {

	var userReq UserLoginRequest

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	_, err := ctx.Request.Cookie(userReq.UserName)

	if err == nil {
		// Cookie already exists, do not set again
		fmt.Println("Cookie already exists for user:", userReq.UserName)
		ctx.JSON(http.StatusOK, "User already logged in:"+userReq.UserName)
		return
	}

	//To ensure username/password has >8 characters (returns bool)
	validatedUserName := util.ValidateUsername(userReq.UserName)
	validatedPassword := util.ValidatePassword(userReq.UserPassword)

	if !validatedUserName {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username format"})
		return
	}

	if !validatedPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password format"})
		return
	}

	//Retrives the credential if does not exist, returns error
	userCred, err := server.store.GetUserByUsername(ctx, userReq.UserName)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		fmt.Println(userCred)
		return
	}

	correctPassword := util.CheckStringHash(userReq.UserPassword, userCred.UserPassword)

	if !correctPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	server.SetCookie(ctx, userCred, time.Hour)
	fmt.Println("Set cookies for:" + userCred.UserName)
	ctx.JSON(http.StatusOK, "User logged in:"+userCred.UserName)
}

func (server *Server) userLogout(ctx *gin.Context) {
	var userReq UserLogoutRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	server.ClearCookie(ctx.Writer, userReq.UserName)

	ctx.JSON(http.StatusOK, "Cookies deleted for:"+userReq.UserName)
}

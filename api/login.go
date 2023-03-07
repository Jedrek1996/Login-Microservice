package api

import (
	db "Microservice-Login/database/sqlc"
	"Microservice-Login/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	UserName     string `json:"user_name" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}

func (server *Server) userLogin(ctx *gin.Context) {

	var userReq UserLoginRequest

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
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

	arg := db.CheckUserCredentialsParams{
		UserName:     userReq.UserName,
		UserPassword: userReq.UserPassword,
	}

	userCred, err := server.store.CheckUserCredentials(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userCred)

}

package api

import (
	db "Microservice-Login/database/sqlc"
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

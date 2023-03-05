package api

import (
	db "Microservice-Login/database/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	UserName     string `json:"user_name" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Mobile       int32  `json:"mobile" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		UserName:     req.UserName,
		UserPassword: req.UserPassword,
		Email:        req.Email,
		Mobile:       req.Mobile,
	}

	account, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

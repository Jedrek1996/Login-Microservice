package api

import (
	db "Microservice-Login/database/sqlc"
	"Microservice-Login/util"
	"log"
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

	validatedPassword := util.ValidatePassword(req.UserPassword) //To ensure password has >8 characters (returns bool)

	if validatedPassword {
		hashedPassword, err := util.HashString(req.UserPassword)

		if err != nil {
			log.Fatalln(err)
			return
		}

		arg := db.CreateUserParams{
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			UserName:     req.UserName,
			UserPassword: hashedPassword,
			Email:        req.Email,
			Mobile:       req.Mobile,
		}

		account, err := server.store.CreateUser(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, account)

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username format"}) // Not >8 characters
		return
	}
}

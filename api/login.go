package api

import (
	"fmt"
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

	fmt.Println("Logging test")

	//Replace with get request
	// account, err := server.store.CreateUser(ctx, arg)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK, userReq)

}

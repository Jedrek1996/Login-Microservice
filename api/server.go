package api

import (
	db "Microservice-Login/database/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Serves http request
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Creates new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/createUser", server.createUser)
	router.POST("/userLogin", server.userLogin)
	router.POST("/userLogout", server.userLogout)
	router.POST("/test", server.AuthCookieMiddleware(), server.TestCookie)

	router.GET("/protected_route", server.AuthCookieMiddleware(), func(c *gin.Context) {
		// This route is protected and can only be accessed by authenticated users
		// If the middleware function returns an unauthorized error, this function will not be executed
		c.JSON(http.StatusOK, gin.H{"message": "This is a protected route"})
	})

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

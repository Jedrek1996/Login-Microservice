package api

import (
	db "Microservice-Login/database/sqlc"

	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Content-Type"}
	config.AllowMethods = []string{"OPTIONS, POST, GET, PUT, DELETE"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.POST("/createUser", server.createUser)
	router.POST("/userLogin", server.userLogin)
	router.POST("/userLogout", server.userLogout)
	router.POST("/getUserData", server.getUserDetail)
	router.POST("/test", server.TestCookie)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

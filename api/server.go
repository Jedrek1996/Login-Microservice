package api

import (
	db "Microservice-Login/database/sqlc"

	"github.com/gin-gonic/gin"
)

// Serves http request
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Creates new http server and setup routing
func newServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/test1", server.createUser)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

package api

import (
	db "github.com/Xebec19/pathshaala-server/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	query  *db.Queries
	router *gin.Engine
}

func NewServer(query *db.Queries) *Server {
	server := &Server{query: query}
	router := gin.Default()

	router.POST("/register", server.createAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

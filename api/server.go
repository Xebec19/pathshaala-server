package api

import (
	"github.com/Xebec19/pathshaala-server/api/auth"
	db "github.com/Xebec19/pathshaala-server/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	query  *db.Queries
	router *gin.Engine
}

func (api *Server) NewServer(query *db.Queries) *Server {
	server := &Server{query: query}
	router := gin.Default()
	auth.Routes(router)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

package api

import (
	"github.com/Xebec19/pathshaala-server/auth"
	db "github.com/Xebec19/pathshaala-server/db/sqlc"
	"github.com/gin-gonic/gin"
)

func NewServer(query *db.Queries) *Server {
	server := &Server{query: query}
	router := gin.Default()
	auth.Routes(router)
	// todo trusted proxies
	// server.router.SetTrustedProxies(nil)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

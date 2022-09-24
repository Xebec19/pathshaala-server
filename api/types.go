package api

import (
	db "github.com/Xebec19/pathshaala-server/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	query  *db.Queries
	router *gin.Engine
}

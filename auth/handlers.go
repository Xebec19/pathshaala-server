package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func register(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "working fine!"})
	return
}

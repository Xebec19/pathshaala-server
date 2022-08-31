package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Register route not working"})
}

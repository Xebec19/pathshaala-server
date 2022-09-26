package auth

import (
	"net/http"

	"github.com/Xebec19/pathshaala-server/utils"
	"github.com/gin-gonic/gin"
)

// register creates a new user account
func register(ctx *gin.Context) {
	var newUser usersParams
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "working fine!"})
	return
}

package auth

import "github.com/gin-gonic/gin"

func Routes(globalRoute *gin.Engine) {
	moduleRoute := globalRoute.Group("/auth")
	moduleRoute.POST("/register", register)
}

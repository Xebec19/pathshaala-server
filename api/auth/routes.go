package auth

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	subRoutes := r.Group("/auth")
	subRoutes.GET("/register", register)
}

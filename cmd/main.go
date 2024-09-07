package main

import (
	"go-class/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// user registration and login interface
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)

	// 需要认证的接口
	auth := r.Group("/")
	auth.Use(api.AuthMiddleware())
	auth.GET("/dashboard", func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to your dashboard, " + username})
	})

	r.Run(":8080")
}

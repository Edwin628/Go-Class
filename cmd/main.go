package main

import (
	"go-class/api"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 配置 CORS，允许所有来源访问（也可以只允许特定来源）
	r.Use(cors.Default())
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

	// WebRTC 信令接口，进行媒体连接的协商
	auth.GET("/webrtc-signal", api.HandleWebRTCSignal)

	r.Run(":8080")
}

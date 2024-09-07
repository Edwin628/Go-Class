package api

import (
	"go-class/sfu"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 创建一个 Upgrader 实例，用于处理 WebSocket 升级
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		// 如果需要跨域支持，可以在这里进行配置
		// 比如检查请求来源，或者直接返回 true 表示允许所有来源
		return true
	},
}

func HandleWebRTCSignal(c *gin.Context) {
	// 升级 HTTP 为 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}

	username := "test-user" // 在实际应用中，你可以通过 JWT 获取用户名

	// 处理 WebRTC 信令
	// peerConnection := sfu.NewPeerConnection()
	// sfu.HandleWebRTCSignal(conn, peerConnection, username)
	sfu.HandleWebRTCSignal(conn, username)
}

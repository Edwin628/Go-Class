package sfu

import (
	"log"

	"github.com/pion/webrtc/v3"
)

// 初始化 SFU
func InitSFU() *webrtc.PeerConnection {
	config := webrtc.Configuration{}

	// 创建 WebRTC 连接
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Fatalf("Failed to create peer connection: %v", err)
	}

	return peerConnection
}

package sfu

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pion/rtp"
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

// // 初始化 SFU
// func InitSFU() {
//     log.Println("SFU initialized")
// }

var (
	mu              sync.Mutex
	peerConnections = make(map[string]*webrtc.PeerConnection) // 保存所有客户端的 PeerConnection
)

// 创建新的 PeerConnection
func NewPeerConnection() *webrtc.PeerConnection {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Fatalf("Failed to create PeerConnection: %v", err)
	}
	return peerConnection
}

// 处理 WebRTC 信令并转发视频流
func HandleWebRTCSignal(conn *websocket.Conn, username string) {
	defer conn.Close()

	// 创建 PeerConnection
	peerConnection := NewPeerConnection()

	// 将这个客户端的 PeerConnection 保存起来
	mu.Lock()
	peerConnections[username] = peerConnection
	mu.Unlock()

	// 处理 ICE 候选
	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate == nil {
			return
		}
		candidateJSON, _ := json.Marshal(candidate.ToJSON())
		conn.WriteMessage(websocket.TextMessage, candidateJSON)
	})

	// 接收到 Track（视频或音频流）时
	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		log.Printf("Received track from user: %s", username)

		// 开始转发音视频流给其他客户端
		forwardStreamToOtherClients(track, username)
	})

	// 处理信令消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		// 解析 SDP 消息
		var sdp webrtc.SessionDescription
		err = json.Unmarshal(message, &sdp)
		if err != nil {
			log.Println("Failed to unmarshal SDP:", err)
			continue
		}

		// 处理远端 SDP
		err = peerConnection.SetRemoteDescription(sdp)
		if err != nil {
			log.Println("Failed to set remote description:", err)
			continue
		}

		// 如果收到 offer，创建 answer
		if sdp.Type == webrtc.SDPTypeOffer {
			answer, err := peerConnection.CreateAnswer(nil)
			if err != nil {
				log.Println("Failed to create answer:", err)
				continue
			}
			err = peerConnection.SetLocalDescription(answer)
			if err != nil {
				log.Println("Failed to set local description:", err)
				continue
			}
			answerJSON, _ := json.Marshal(answer)
			conn.WriteMessage(websocket.TextMessage, answerJSON)
		}
	}
}

// 转发流到其他客户端
func forwardStreamToOtherClients(track *webrtc.TrackRemote, senderUsername string) {
	mu.Lock()
	defer mu.Unlock()

	// 遍历所有连接的 PeerConnection，将音视频流转发给其他客户端
	for username, pc := range peerConnections {
		// 不转发给发送者自己
		if username == senderUsername {
			continue
		}

		// 获取媒体格式信息
		trackCodec := track.Codec()

		// 为其他客户端创建新的 Track
		localTrack, err := webrtc.NewTrackLocalStaticRTP(
			trackCodec.RTPCodecCapability,
			track.ID(),
			track.StreamID(),
		)
		if err != nil {
			log.Printf("Failed to create local track: %v", err)
			continue
		}

		// 将 Track 添加到 PeerConnection 中
		_, err = pc.AddTrack(localTrack)
		if err != nil {
			log.Printf("Failed to add track to peer: %v", err)
			continue
		}

		// 将数据流写入新的 Track 中
		rtpBuf := make([]byte, 1400)
		for {
			// 从远程 Track 中读取 RTP 包
			i, _, readErr := track.Read(rtpBuf)
			if readErr != nil {
				log.Println("Error reading track:", readErr)
				break
			}

			// 解析 RTP 包
			rtpPacket := &rtp.Packet{}
			if err := rtpPacket.Unmarshal(rtpBuf[:i]); err != nil {
				log.Printf("Failed to unmarshal RTP packet: %v", err)
				break
			}

			// 将 RTP 包写入本地 Track
			if writeErr := localTrack.WriteRTP(rtpPacket); writeErr != nil {
				log.Printf("Error writing RTP packet to local track: %v", writeErr)
				break
			}
		}
	}
}

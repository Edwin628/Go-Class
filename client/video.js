// 检查用户是否已经登录
const token = localStorage.getItem('token');
if (!token) {
    window.location.href = 'login.html'; // 未登录则跳转到登录页面
}

// 退出功能：清除 JWT 令牌并返回登录页面
document.getElementById('logoutButton').addEventListener('click', () => {
    localStorage.removeItem('token');  // 清除 JWT 令牌
    window.location.href = 'login.html';  // 重定向到登录页面
});

const localVideo = document.getElementById('localVideo');
const remoteVideo = document.getElementById('remoteVideo');

// 获取本地音视频流
navigator.mediaDevices.getUserMedia({ video: true, audio: true })
    .then(stream => {
        localVideo.srcObject = stream;
        stream.getTracks().forEach(track => peerConnection.addTrack(track, stream));
    });

const socket = new WebSocket('ws://localhost:8080/webrtc-signal');
const peerConnection = new RTCPeerConnection({
    iceServers: [{ urls: 'stun:stun.l.google.com:19302' }]
});

// WebSocket 信令
socket.onopen = () => {
    // WebSocket 连接成功后可以发送消息
    console.log("WebSocket connection established");

    peerConnection.createOffer().then(offer => {
        return peerConnection.setLocalDescription(offer);
    }).then(() => {
        socket.send(JSON.stringify(peerConnection.localDescription)); // 发送 SDP
    });
};

socket.onmessage = (message) => {
    const data = JSON.parse(message.data);
    if (data.sdp) {
        peerConnection.setRemoteDescription(new RTCSessionDescription(data.sdp));
    } else if (data.candidate) {
        peerConnection.addIceCandidate(new RTCIceCandidate(data.candidate));
    }
};

socket.onerror = (error) => {
    console.error("WebSocket error: ", error);
};

// 接收远程视频流
peerConnection.ontrack = (event) => {
    remoteVideo.srcObject = event.streams[0];
};

// ICE 候选
peerConnection.onicecandidate = (event) => {
    if (event.candidate) {
        socket.send(JSON.stringify(event.candidate));
    }
};

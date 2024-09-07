# Go-Class - 多人视频聊天应用

## 项目简介

Go-Class是一个使用 Go 语言开发的多人视频聊天应用，支持通过 SFU（Selective Forwarding Unit）架构进行实时视频传输，并提供了用户注册、登录、房间管理、录播等功能。项目采用了 WebRTC 技术实现低延迟的视频聊天功能，后台使用 JWT 进行身份验证，并提供了简洁的 API 接口。

## 目录结构

```
go-class/
│
├── cmd/
│   └── main.go               # 主程序入口
├── api/
│   ├── auth.go               # 用户认证逻辑
│   ├── rooms.go              # 房间管理 API
│   └── recordings.go         # 录播管理
├── sfu/
│   ├── sfu.go                # SFU 初始化
│   └── recorder.go           # 录播处理
├── config/
│   └── config.go             # 项目配置
├── models/
│   └── user.go               # 用户模型与数据库交互
├── go.mod                    # Go 模块文件
└── README.md                 # 项目文档
```

## 功能概述

1. **用户注册和登录**：用户可以通过提供用户名和密码进行注册和登录，登录后将获得 JWT token 用于后续的 API 访问。
2. **视频房间管理**：允许用户创建和加入视频房间。管理员可以删除房间。
3. **多人视频聊天**：基于 WebRTC 技术，使用 SFU 架构实现多人实时视频通话。
4. **录播功能**：提供录制视频会议的功能，录制的视频可以保存并在后台管理。
5. **后台管理**：管理员可以查看并管理房间、用户和录制视频。

## 安装与运行

### 环境要求

- [Go 1.18+](https://golang.org/dl/)
- [FFmpeg](https://ffmpeg.org/download.html)（用于录播功能）
- Nginx + RTMP（可选，用于流媒体转发）

### 本地运行步骤

1. 克隆项目到本地：

   ```bash
   git clone https://github.com/your-username/streaming-app.git
   cd streaming-app
   ```

2. 初始化 Go 模块：

   ```bash
   go mod tidy
   ```

3. 运行项目：

   ```bash
   go run cmd/main.go
   ```

4. 访问 `http://localhost:8080`，通过 API 测试用户注册和登录功能。

### Docker 部署（可选）

如果你希望使用 Docker 部署项目，确保已经安装 Docker，然后可以执行以下步骤：

1. 构建 Docker 镜像：

   ```bash
   docker build -t streaming-app .
   ```

2. 运行 Docker 容器：

   ```bash
   docker run -p 8080:8080 streaming-app
   ```

### 测试 API

你可以使用 `curl` 或 Postman 进行 API 调试，以下是一些示例 API 请求。

#### 用户注册

```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"username": "user1", "password": "pass123"}'
```

#### 用户登录

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username": "user1", "password": "pass123"}'
```

成功登录后会返回一个 JWT token，用于后续的 API 请求。

#### 创建房间

```bash
curl -X POST http://localhost:8080/admin/create-room \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_token>" \
-d '{"name": "room1", "host": "user1"}'
```

#### 获取房间列表

```bash
curl -X GET http://localhost:8080/admin/rooms \
-H "Authorization: Bearer <your_token>"
```

## 配置

项目配置文件 `config/config.go` 中可以修改应用的配置参数，比如数据库连接、WebRTC 配置、日志级别等。

## 未来改进

- **完善前端界面**：目前项目主要以后端 API 为主，可以增加前端界面（使用 React/Vue.js）来提供用户友好的交互界面。
- **负载均衡和高并发优化**：可以使用 Nginx 或其他负载均衡工具对高并发进行优化。
- **实时聊天和屏幕共享**：增加多人视频会议中的实时文本聊天和屏幕共享功能。

## 贡献

欢迎你为这个项目贡献代码！请先提交 Issue，描述你的建议或 bug 报告。提交 Pull Request 前，请确保你的代码通过了基本测试。

## 许可证

本项目遵循 MIT 许可证。详细信息请查阅 `LICENSE` 文件。

---
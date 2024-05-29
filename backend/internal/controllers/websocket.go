package controllers

import (
	"backend/internal/models"
	"backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 设置WebSocket升级器，允许跨域
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[uint]*websocket.Conn)

func WebSocketMessage(c *gin.Context) {
	// 从请求中获取Token
	token := c.Query("token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		log.Default().Println("没有token")
		return
	}
	userID, err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		log.Default().Println("token验证失败", err.Error())
		return
	}

	// 将HTTP链接升级为WebSocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Default().Println("Failed to upgrade connection to WebSocket")
		return
	}

	// 从请求中获取用户ID，设置为消息发送者
	clients[userID] = ws

	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Default().Println("Failed to close WebSocket connection")
		}
	}(ws)
	defer delete(clients, userID)

	for {
		var msgGet models.MessageRequestBody
		err := ws.ReadJSON(&msgGet)
		if err != nil {
			respondWebSocket(ws, 1, "读取信息失败", nil)
			log.Default().Println("Failed to read message from WebSocket, Breaking loop, err:", err)
			break
		}

		// 如果是心跳包，直接返回
		if msgGet.Type == "ping" {
			respondWebSocket(ws, 2, "ok", nil)
			continue
		}

		// 检查用户是否在聊天室中
		if !userRoomRepo.IfUserInRoom(userID, msgGet.RoomID) {
			respondWebSocket(ws, 1, "当前用户已不在聊天室中", nil)
			continue
		}
		
		message := models.Message{
			RoomID:   msgGet.RoomID,
			SenderID: userID,
			Type:     msgGet.Type,
			Content:  msgGet.Content,
		}

		newMessage, err := messageRepo.CreateMessage(message)
		if err != nil {
			respondWebSocket(ws, 1, "创建新信息失败", nil)
			continue
		}

		// 获取聊天室内的所有用户
		roomUsers, err := userRoomRepo.GetRoomUsers(msgGet.RoomID)
		if err != nil {
			respondWebSocket(ws, 1, "获取房间内用户失败", nil)
			continue
		}

		senderInfo, err := userInfoRepository.GetUserInfoByID(userID)
		if err != nil {
			respondWebSocket(ws, 1, "获取发件人名称失败", nil)
			continue
		}

		msgSend := models.MessageResponseBody{
			ID:           newMessage.ID,
			SendTime:     newMessage.CreatedAt.Format("2006-01-02 15:04:05"),
			RoomID:       newMessage.RoomID,
			SenderID:     newMessage.SenderID,
			SenderName:   senderInfo.Username,
			SenderAvatar: senderInfo.Avatar,
			Type:         newMessage.Type,
			Content:      newMessage.Content,
		}

		// 发送消息给聊天室内的所有在线用户
		for _, user := range roomUsers {
			if client, ok := clients[user.ID]; ok {
				respondWebSocket(client, 0, "新消息", msgSend)
			}
		}
	}
}

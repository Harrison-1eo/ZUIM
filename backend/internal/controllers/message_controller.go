// controllers/message_controller.go

package controllers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
)

var messageRepo *repositories.MessageRepository = repositories.NewMessageRepository()
var fileRepo *repositories.FileRepository = repositories.NewFileRepository()

// SendMessage POST方法 发送消息
func SendMessage(c *gin.Context) {
	type RequestBody struct {
		RoomID  uint   `json:"room_id"`
		Type    string `json:"type"`
		Content string `json:"content"`
	}
	var msg RequestBody
	if err := c.ShouldBindJSON(&msg); err != nil {
		respond(c, 1, "Invalid message data", nil)
		return
	}

	// 从请求中获取用户ID，设置为消息发送者
	userID := c.MustGet("userID").(uint)
	message := models.Message{
		RoomID:   msg.RoomID,
		SenderID: userID,
		Type:     msg.Type,
		Content:  msg.Content,
	}

	newMessage, err := messageRepo.CreateMessage(message)
	if err != nil {
		respond(c, 1, "Failed to send message", nil)
		return
	}

	type ResponseType struct {
		ID       uint   `json:"ID"`
		SendTime string `json:"send_time"`
	}

	respond(c, 0, "Message sent successfully",
		ResponseType{ID: newMessage.ID,
			SendTime: newMessage.CreatedAt.Format("2006-01-02 15:04:05")})
}

// 获取聊天室消息
func GetMessages(c *gin.Context) {
	type RequestBody struct {
		RoomID        uint `json:"room_id"`
		LastMessageID uint `json:"last_message_id"`
		Limit         int  `json:"limit"`
	}
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		respond(c, 1, "Invalid query message data", nil)
		return
	}
	messages, err := messageRepo.GetMessagesByRoomID(body.RoomID, body.LastMessageID, body.Limit)
	if err != nil {
		respond(c, 1, "Failed to retrieve messages", nil)
		return
	}

	sm := make([]models.MessageResponseBody, len(messages))
	for i, m := range messages {
		senderName, err := userRepository.GetUsernameByID(m.SenderID)
		if err != nil {
			respond(c, 1, "Failed to retrieve messages", nil)
			return
		}
		sm[i] = models.MessageResponseBody{
			ID:         m.ID,
			SendTime:   m.CreatedAt.Format("2006-01-02 15:04:05"),
			RoomID:     m.RoomID,
			SenderID:   m.SenderID,
			SenderName: senderName,
			Type:       m.Type,
			Content:    m.Content,
		}
	}

	respond(c, 0, "历史消息获取成功", sm)
}

// 上传文件
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		respond(c, 1, "上传文件失败", nil)
		return
	}

	userID := c.MustGet("userID").(uint)
	roomID := c.PostForm("room_id")
	println("roomID:", roomID)
	if roomID == "" {
		respond(c, 1, "无效的房间roomid not found", nil)
		return
	}
	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil {
		respond(c, 1, "无效的房间 damn", nil)
		return
	}

	// 保存文件到本地，保存路径为 /static/files/{room_id}/
	fileName := utils.GenerateFilename(file.Filename)
	filePath := "static/files/" + fileName
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		respond(c, 1, "上传文件失败", nil)
		return
	}

	// 保存文件信息到数据库
	fileInfo := models.File{
		FileName: file.Filename,
		FileType: path.Ext(file.Filename),
		FilePath: filePath,
		RoomID:   uint(roomIDInt),
		SenderID: userID,
	}
	_, err = fileRepo.CreateFile(fileInfo)
	if err != nil {
		respond(c, 1, "上传文件失败", nil)
		return
	}

	type ResponseType struct {
		CreateTime string `json:"create_time"`
		FileName   string `json:"file_name"`
		FileType   string `json:"file_type"`
		FileURL    string `json:"file_url"`
		RoomID     uint   `json:"room_id"`
	}

	respond(c, 0, "上传文件成功",
		ResponseType{
			CreateTime: fileInfo.CreatedAt.Format("2006-01-02 15:04:05"),
			FileName:   fileInfo.FileName,
			FileType:   fileInfo.FileType,
			FileURL:    "/" + fileInfo.FilePath,
			RoomID:     fileInfo.RoomID,
		})
}

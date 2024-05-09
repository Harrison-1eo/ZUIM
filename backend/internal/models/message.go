// models/message.go

package models

import (
	"gorm.io/gorm"
)

// 消息模型
type Message struct {
	gorm.Model
	SenderID uint   `json:"sender_id" gorm:"not null"`
	RoomID   uint   `json:"room_id" gorm:"not null"`
	Type     string `json:"type" gorm:"not null"`
	Content  string `json:"content" gorm:"type:text"`
}

type MessageRequestBody struct {
	RoomID  uint   `json:"room_id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type MessageResponseBody struct {
	ID         uint   `json:"ID"`
	SendTime   string `json:"send_time"`
	RoomID     uint   `json:"room_id"`
	SenderID   uint   `json:"sender_id"`
	SenderName string `json:"sender_name"`
	Type       string `json:"type"`
	Content    string `json:"content"`
}

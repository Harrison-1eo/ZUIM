package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	FileName string `json:"file_name" gorm:"not null"`
	FileType string `json:"file_type" gorm:"not null"`
	FilePath string `json:"file_path" gorm:"not null"`
	RoomID   uint   `json:"room_id" gorm:"not null"`
	SenderID uint   `json:"sender_id" gorm:"not null"`
}

// models/user_room.go

package models

import "gorm.io/gorm"

type UserRoom struct {
	gorm.Model
	RoomID uint   `json:"room_id" gorm:"not null;index"`
	UserID uint   `json:"user_id" gorm:"not null;index"`
	Role   string `json:"role"` // 如 "member", "admin", "owner"等
}

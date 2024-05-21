package models

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	UserID   uint   `gorm:"index" json:"user_id"`
	FriendID uint   `gorm:"index" json:"friend_id"`
	Remark   string `json:"remark"`
}

// models/user.go

package models

import (
	"gorm.io/gorm"
	"time"
)

// 用户模型
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type OnlineUsers struct {
	// CREATE TABLE online_users (
	//    user_id INT NOT NULL,
	//    login_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	//    PRIMARY KEY (user_id)
	//) ENGINE=MEMORY;
	UserID    uint      `json:"user_id" gorm:"primaryKey;not null"`
	LoginTime time.Time `json:"login_time" gorm:"not null"`
}

// models/room.go

package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
}

// repositories/room_repository.go

package repositories

import (
	"IMTest/internal/models"
)

type RoomRepository struct {
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{}
}

// 创建聊天室
func (repo *RoomRepository) CreateRoom(room models.Room) (*models.Room, error) {
	if err := db.Model(models.Room{}).Create(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

// 获取聊天室信息
func (repo *RoomRepository) GetRoom(id uint) (*models.Room, error) {
	var room models.Room
	if err := db.Model(models.Room{}).First(&room, id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

// repositories/room_repository.go

package repositories

import (
	"backend/internal/models"
)

type RoomRepository struct {
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{}
}

// CreateRoom 创建聊天室
func (repo *RoomRepository) CreateRoom(room models.Room) (*models.Room, error) {
	if err := db.Model(models.Room{}).Create(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

// DeleteRoom 删除聊天室
func (repo *RoomRepository) DeleteRoom(id uint) error {
	if err := db.Delete(&models.Room{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetRoom 获取聊天室信息
func (repo *RoomRepository) GetRoom(id uint) (*models.Room, error) {
	var room models.Room
	if err := db.Model(models.Room{}).First(&room, id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

// GetRoomWithDeleted 获取聊天室信息，包含被删除的聊天室
func (repo *RoomRepository) GetRoomWithDeleted(id uint) (*models.Room, error) {
	var room models.Room
	if err := db.Model(models.Room{}).Unscoped().First(&room, id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

// UpdateRoom 修改聊天室信息
func (repo *RoomRepository) UpdateRoom(room models.Room) (*models.Room, error) {
	if err := db.Model(models.Room{}).Where("id = ?", room.ID).Save(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

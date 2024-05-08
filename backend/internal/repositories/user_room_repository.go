// repositories/user_room_repository.go

package repositories

import (
	"backend/internal/models"
)

type UserRoomRepository struct {
}

func NewUserRoomRepository() *UserRoomRepository {
	return &UserRoomRepository{}
}

// 添加用户到聊天室
func (repo *UserRoomRepository) AddUserToRoom(userID, roomID uint, role string) error {
	userRoom := models.UserRoom{
		RoomID: roomID,
		UserID: userID,
		Role:   role,
	}
	if err := db.Model(models.UserRoom{}).Create(&userRoom).Error; err != nil {
		return err
	}
	return nil
}

// 获取聊天室所有用户，user和userRoom表关联查询
func (repo *UserRoomRepository) GetRoomUsers(roomID uint) ([]models.User, error) {
	var users []models.User
	if err := db.Model(models.UserRoom{}).
		Select("users.ID, users.Username").
		Joins("left join users on users.id = user_rooms.user_id").
		Where("user_rooms.room_id = ?", roomID).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// 判断用户是否在聊天室中
func (repo *UserRoomRepository) IfUserInRoom(userID, roomID uint) bool {
	var count int64
	db.Model(models.UserRoom{}).Where("user_id = ? AND room_id = ?", userID, roomID).Count(&count)
	return count > 0
}

// GetRoomsByUserID 获取某个用户的聊天室列表，room和userRoom表关联查询
func (repo *UserRoomRepository) GetRoomsByUserID(userID uint) ([]models.Room, error) {
	var rooms []models.Room
	if err := db.Model(models.UserRoom{}).
		Select("rooms.id, rooms.name, rooms.description").
		Joins("left join rooms on rooms.id = user_rooms.room_id").
		Where("user_rooms.user_id = ?", userID).
		Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

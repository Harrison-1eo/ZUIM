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

func (repo *UserRoomRepository) GetRoomUserInfos(roomID uint) ([]models.UserInfo, error) {
	var userInfos []models.UserInfo
	if err := db.Model(models.UserRoom{}).
		Select("user_infos.*").
		Joins("left join user_infos on user_infos.user_id = user_rooms.user_id").
		Where("user_rooms.room_id = ?", roomID).
		Find(&userInfos).Error; err != nil {
		return nil, err
	}
	return userInfos, nil
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

func (repo *UserRoomRepository) IfUserIsOwner(userId uint, roomId uint) bool {
	var userRoom models.UserRoom
	res := db.Model(models.UserRoom{}).Where("UserID = ? AND RoomID = ? AND Role = ?", userId, roomId, "owner")
	if res.First(&userRoom).Error != nil {
		return false
	}
	return true
}

func (repo *UserRoomRepository) DeleteRoom(roomID uint) error {
	if err := db.Delete(&models.UserRoom{}, "room_id = ?", roomID).Error; err != nil {
		return err
	}
	return nil
}

// GetCommonRooms 获取和某个人一起在的聊天室
func (repo *UserRoomRepository) GetCommonRooms(userID1, userID2 uint) ([]models.Room, error) {
	var rooms []models.Room
	if err := db.Model(models.UserRoom{}).
		Select("rooms.id, rooms.name, rooms.description").
		Joins("left join rooms on rooms.id = user_rooms.room_id").
		Where("user_rooms.user_id = ? OR user_rooms.user_id = ?", userID1, userID2).
		Group("rooms.id").
		Having("count(rooms.id) = 2").
		Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

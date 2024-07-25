package repositories

import (
	"backend/internal/models"
	"time"
)

type OnlineUserRepository struct {
}

// NewOnlineUserRepository 初始化用户仓库
func NewOnlineUserRepository() *OnlineUserRepository {
	return &OnlineUserRepository{}
}

// SetOnline 更新登录状态
func (repo *OnlineUserRepository) SetOnline(userID uint) error {
	// 删除旧记录
	db.Delete(&models.OnlineUsers{}, "user_id = ?", userID)

	// 插入新记录
	if err := db.Create(&models.OnlineUsers{
		UserID:    userID,
		LoginTime: time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

// SetOffline 更新登出状态
func (repo *OnlineUserRepository) SetOffline(userID uint) error {
	return db.Delete(&models.OnlineUsers{}, "user_id = ?", userID).Error
}

// CountOnlineUsers 统计在线人数
func (repo *OnlineUserRepository) CountOnlineUsers() (int, error) {
	var count int64
	if err := db.Model(&models.OnlineUsers{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

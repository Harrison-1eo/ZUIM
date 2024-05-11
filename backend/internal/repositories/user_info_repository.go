package repositories

import (
	"backend/internal/models"
	"errors"
	"gorm.io/gorm"
)

type UserInfoRepository struct{}

func NewUserInfoRepository() *UserInfoRepository {
	return &UserInfoRepository{}
}

func (repo *UserInfoRepository) GetUserInfoByID(id uint) (*models.UserInfo, error) {
	var userInfo models.UserInfo
	if err := db.Model(models.UserInfo{}).Where("user_id = ?", id).First(&userInfo).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (repo *UserInfoRepository) CreateUserInfo(userInfo models.UserInfo) (*models.UserInfo, error) {
	if err := db.Model(models.UserInfo{}).Create(&userInfo).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (repo *UserInfoRepository) UpdateUserInfo(userInfo models.UserInfo) (*models.UserInfo, error) {
	// 首先检查用户信息是否存在，若不存在则创建
	if _, err := repo.GetUserInfoByID(userInfo.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repo.CreateUserInfo(userInfo)
		}
	}
	if err := db.Model(models.UserInfo{}).Save(&userInfo).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}

package repositories

import "backend/internal/models"

func (repo *UserRepository) GetAllUsers() ([]models.UserInfo, error) {
	// 使用联合查询获取所有用户信息
	var userInfos []models.UserInfo
	if err := db.Model(models.UserInfo{}).Find(&userInfos).Error; err != nil {
		return nil, err
	}
	return userInfos, nil
}

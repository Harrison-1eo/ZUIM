// repositories/user_repository.go

package repositories

import "backend/internal/models"

type UserRepository struct {
}

// 初始化用户仓库
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// 检查用户名是否已存在
func (repo *UserRepository) UserExists(username string) bool {
	var user models.User
	println("username : ", username)
	res := db.Model(models.User{}).Where("Username = ?", username)
	if res.First(&user).Error != nil {
		return false
	}
	return true
}

// 创建用户
func (repo *UserRepository) CreateUser(user models.User) (*models.User, error) {
	if err := db.Model(models.User{}).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 验证用户登录信息
func (repo *UserRepository) AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	if err := db.Model(models.User{}).Where("Username = ? AND Password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

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
	//println("username : ", username)
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

// 根据ID获取用户
func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := db.Model(models.User{}).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 根据ID获取用户名
func (repo *UserRepository) GetUsernameByID(id uint) (string, error) {
	var user models.User
	if err := db.Model(models.User{}).First(&user, id).Error; err != nil {
		return "", err
	}
	return user.Username, nil
}

func (repo *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := db.Model(models.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	//println("username : ", username)
	//println("user : ", user.ID)
	return &user, nil
}

func (repo *UserRepository) UpdatePassword(id uint, password string) error {
	return db.Model(models.User{}).Where("id = ?", id).Update("password", password).Error
}

func (repo *UserRepository) UpdateUser(user models.User) (*models.User, error) {
	if err := db.Model(models.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

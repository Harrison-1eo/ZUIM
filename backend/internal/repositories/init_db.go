// internal/config/init_db.go

package repositories

import (
	"backend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// 初始化数据库连接并迁移数据库模式
func InitDB() {
	// 连接SQLite数据库
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		// 打印SQL语句
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("internal/config/init_db.go: failed to connect database >>> " + err.Error())
	}

	// 迁移数据库模式
	err = db.AutoMigrate(&models.User{}, &models.Room{}, &models.UserRoom{}, &models.Message{})
	if err != nil {
		panic("internal/config/init_db.go: failed to migrate database >>> " + err.Error())
	}

	// 向用户表中插入一条记录用于测试
	db.Create(&models.User{Username: "admin", Password: "admin"})

	// 从数据库中获取第一条用户记录
	var user models.User
	db.First(&user)
	println("internal/config/init_db.go: first user >>> ", user.Username, user.Password)

}

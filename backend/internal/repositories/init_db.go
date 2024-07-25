// internal/config/init_db.go

package repositories

import (
	"backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// InitDBSqlite 初始化数据库连接并迁移数据库模式
func InitDBSqlite() {
	// 连接SQLite数据库
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		// 打印SQL语句
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("internal/config/init_db.go: failed to connect sqlite database >>> " + err.Error())
	}

	// 迁移数据库模式
	err = db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.UserRoom{},
		&models.Message{},
		&models.File{},
		&models.UserInfo{},
	)
	if err != nil {
		panic("internal/config/init_db.go: failed to migrate sqlite database >>> " + err.Error())
	}

	// 向用户表中插入一条记录用于测试
	//db.Create(&models.User{Username: "admin", Password: "admin"})

	// 从数据库中获取第一条用户记录
	//var user models.User
	//db.First(&user)
	//println("internal/config/init_db.go: first user >>> ", user.Username, user.Password)

	// 对用户表进行查询，每个没有Userinfo的用户插入一条Userinfo记录
	//var users []models.User
	//db.Find(&users)
	//for _, user := range users {
	//	var userInfo models.UserInfo
	//	db.Where("user_id = ?", user.ID).First(&userInfo)
	//	if userInfo.UserID == 0 {
	//		db.Create(&models.UserInfo{
	//			UserID:   user.ID,
	//			Username: user.Username,
	//		})
	//	}
	//}
}

func InitDBMySQL() {
	// 连接MySQL数据库
	var err error
	db, err = gorm.Open(mysql.Open("imdsbMaster:imdsbPassword@tcp(127.0.0.1:3306)/imdsb?parseTime=true"), &gorm.Config{
		// 打印SQL语句
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("internal/config/init_db.go: failed to connect mysql database >>> " + err.Error())
	}

	// 迁移数据库模式
	// 迁移数据库模式
	err = db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.UserRoom{},
		&models.Message{},
		&models.File{},
		&models.UserInfo{},
		&models.OnlineUsers{},
	)
	if err != nil {
		panic("internal/config/init_db.go: failed to migrate mysql database >>> " + err.Error())
	}

}

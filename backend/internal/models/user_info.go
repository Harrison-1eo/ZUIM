package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	//User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 指定外键
	UserID    uint   `json:"user_id"   gorm:"index"`    // 外键, 与User表的ID关联
	Username  string `json:"username" gorm:"unique"`    // 用户名
	NikeName  string `json:"nike_name"`                 // 用户全名
	Avatar    string `json:"avatar"`                    // 用户头像
	Sexuality string `json:"sexuality"`                 // 性别
	Year      uint   `json:"year"`                      // 出生年份
	Month     uint   `json:"month"`                     // 出生月份
	Day       uint   `json:"day"`                       // 出生日期
	Country   string `json:"country"`                   // 国家
	Province  string `json:"province"`                  // 省份
	City      string `json:"city"`                      // 城市
	Email     string `json:"email" gorm:"unique,index"` // 电子邮件
}

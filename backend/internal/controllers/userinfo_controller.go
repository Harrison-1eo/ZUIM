package controllers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var userInfoRepository = repositories.NewUserInfoRepository()

// GetMyInfo GET方法 获取用户信息
func GetMyInfo(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	userInfo, err := userInfoRepository.GetUserInfoByID(userID)
	if err != nil {
		respond(c, 1, "获取用户信息失败", nil)
		return
	}
	respond(c, 0, "获取用户信息成功", userInfo)
}

// UpdateUserInfo POST方法 更新用户信息
func UpdateUserInfo(c *gin.Context) {
	var userInfo models.UserInfo
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		respond(c, 1, "无效的请求数据", nil)
		return
	}

	userID := c.MustGet("userID").(uint)
	if userInfo.UserID != userID {
		respond(c, 1, "无权操作", nil)
		return
	}

	updatedUserInfo, err := userInfoRepository.UpdateUserInfo(userInfo)
	if err != nil {
		respond(c, 1, "更新用户信息失败", nil)
		return
	}
	respond(c, 0, "更新用户信息成功", updatedUserInfo)
}

// GetUserInfo GET方法 获取指定用户信息
func GetUserInfo(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		respond(c, 1, "无效的用户ID", nil)
		return
	}

	userInfo, err := userInfoRepository.GetUserInfoByID(uint(userID))
	if err != nil {
		respond(c, 1, "获取用户信息失败", nil)
		return
	}

	type ResponseType struct {
		ID        uint   `json:"user_id"`
		Username  string `json:"username"`
		Avatar    string `json:"avatar"`
		Sexuality string `json:"sexuality"`
		Year      uint   `json:"year"`
		Month     uint   `json:"month"`
		Day       uint   `json:"day"`
		Country   string `json:"country"`
		Province  string `json:"province"`
		City      string `json:"city"`
	}
	respond(c, 0, "获取用户信息成功", ResponseType{
		ID:        userInfo.UserID,
		Username:  userInfo.Username,
		Avatar:    userInfo.Avatar,
		Sexuality: userInfo.Sexuality,
		Year:      userInfo.Year,
		Month:     userInfo.Month,
		Day:       userInfo.Day,
		Country:   userInfo.Country,
		Province:  userInfo.Province,
		City:      userInfo.City,
	})

}

func UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		respond(c, 1, "上传文件失败", nil)
		return
	}

	// 检查是否是图片，Content-Type只能是image/*
	if file.Header.Get("Content-Type")[:5] != "image" {
		respond(c, 1, "上传文件不是图片", nil)
		return
	}

	userID := c.MustGet("userID").(uint)
	fileName := utils.GenerateFilename(file.Filename)
	filePath := "static/avatars/" + fileName
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		respond(c, 1, "上传文件失败", nil)
		return
	}

	updatedUserInfo, err := userInfoRepository.GetUserInfoByID(userID)
	if err != nil {
		respond(c, 1, "获取用户信息失败", nil)
		return
	}
	if updatedUserInfo.Avatar != "" {
		err = utils.DeleteFile(updatedUserInfo.Avatar)
		if err != nil {
			respond(c, 1, "删除旧头像失败", nil)
			return
		}
	}

	updatedUserInfo.Avatar = "/" + filePath
	_, err = userInfoRepository.UpdateUserInfo(*updatedUserInfo)

	if err != nil {
		respond(c, 1, "更新用户信息失败", nil)
		return
	}

	respond(c, 0, "更新用户信息成功", updatedUserInfo)
}

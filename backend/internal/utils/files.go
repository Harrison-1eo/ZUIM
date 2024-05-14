package utils

import (
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

// GenerateFilename 生成唯一的文件名
func GenerateFilename(filename string) string {
	// 生成时间戳+6位随机数+文件后缀
	return strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000) + path.Ext(filename)
}

// DeleteFile 删除static目录下的文件
func DeleteFile(filePath string) error {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}
	// 检查是否是static目录下的文件
	if path.Dir(filePath) != "static" {
		return nil
	}
	return os.Remove(filePath)
}

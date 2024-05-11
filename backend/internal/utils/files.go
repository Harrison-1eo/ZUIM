package utils

import (
	"math/rand"
	"path"
	"strconv"
	"time"
)

// 生成唯一的文件名
func GenerateFilename(filename string) string {
	// 生成时间戳+6位随机数+文件后缀
	return strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000) + path.Ext(filename)
}

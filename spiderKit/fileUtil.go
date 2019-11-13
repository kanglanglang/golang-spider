package spiderKit

import (
	"strconv"
	"time"
)

// GetRandomFileName 生成时间戳_随机数文件名
func GetRandomFileName() string {
	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	randomNum := strconv.Itoa(GetRandomInt(100, 1000))
	return timestamp + "_" + randomNum
}

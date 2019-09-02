package utils

import "time"

// 获取当前时间的Unix时间戳
func GetCurrentTime() int64 {
	return time.Now().Unix()
}

// 获取当前毫秒级时间
func GetCurrentMilliTime() int64 {
	return time.Now().UnixNano() / 1000000
}

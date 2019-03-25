package randtool

import (
	"math/rand"
)

const (
	std32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	std64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

// 获取指定长度随机32个字符的字符串
func Rand32String(length uint8) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = std32[rand.Intn(32)]
	}
	return string(b)
}

// 获取指定长度随机64个字符的字符串
func Rand64String(length uint8) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = std64[rand.Intn(32)]
	}
	return string(b)
}

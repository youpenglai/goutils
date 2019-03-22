package randtool

import (
	"math/rand"
)

// 获取2个数之间的随机数
func RandInt64(min, max int64) int64 {
	if min >= max {
		return max
	}
	return rand.Int63n(max-min) + min
}
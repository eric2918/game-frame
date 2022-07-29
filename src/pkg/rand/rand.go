package rand

import (
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString 随机字符串
// size:随机位数
// kind:随机类型(0:纯数字, 1:小写字母, 2:大写字母, 3:数字、大小写字母)
func RandomString(size int, kind int) string {
	kinds, bytes := [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	for i := 0; i < size; i++ {
		if isAll {
			kind = rand.Intn(3)
		}
		scope, base := kinds[kind][0], kinds[kind][1]
		bytes[i] = uint8(base + rand.Intn(scope))
	}
	return string(bytes)
}

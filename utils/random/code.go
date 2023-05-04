package random

import (
	"math/rand"
	"strconv"
	"time"
)

// Code 随机生成一个length位数的验证码
func Code(length int) string {
	rand.Seed(time.Now().UnixNano())
	max := 1
	for i := 0; i < length; i++ {
		max *= 10
	}
	code := rand.Intn(max)
	str := strconv.Itoa(code)
	zeroNum := length - len(str)
	for i := 0; i < zeroNum; i++ {
		str = "0" + str
	}
	return str
}

package nonce

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// 获取随机数
func getRandStr(seek string, l int) string {
	seekLen := len(seek)
	str := ""
	for i := 0; i < l; i++ {
		key := r.Intn(seekLen)
		str += seek[key : key+1]
	}
	return str
}

// 随机字符串
func Str(l int) string {
	seek := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return getRandStr(seek, l)
}

// 纯数字随机字符串
func Num(l int) string {
	seek := "0123456789"
	return getRandStr(seek, l)
}

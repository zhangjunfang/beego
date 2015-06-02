// 产生随机数或者随机字符串
package util

import (
	"math/rand"
	"time"
)

// 获取范围为[0,max)，类型为int的随机整数
func RandInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

// 获取范围为[0.0, 1.0)，类型为float32的随机小数
func RandFloat32() float32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float32()
}

// 获取范围为[0.0, 1.0)，类型为float64的随机小数
func RandFloat64() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float64()
}

// 获取范围为[0,max)，数量为max，类型为int的随机整数slice
func RandPerm(max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Perm(max)
}

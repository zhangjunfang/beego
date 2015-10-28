package util

import (
	"encoding/hex"
	"strconv"
	"strings"
)

func init() {

}

const ()

/*
函数功能说明：
1.字符串s 首先截取prefix和suffix 后，使用delimiter分割字符串
2.
3.返回 字符串切片
4.
*/
func SplitString(s, delimiter, prefix, suffix string) []string {

	if strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
	}
	if strings.HasSuffix(s, suffix) {
		var length = len(s)
		s = s[:length-1]
	}
	var arrays []string = strings.Split(s, delimiter)
	return arrays
}

/*
函数功能说明：
1.在字符串s 截取首部的prefix 以及尾部的prefix字符串后 返回截取的新字符串
2.
3.
4.
*/
func SubString(s, prefix, suffix string) string {
	if strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
	}
	if strings.HasSuffix(s, suffix) {
		var length = len(s)
		s = s[:length-1]
	}
	return s
}

/*
函数功能介绍：
1. 把字符串转换为字节数组
2. 把字节数组转换为十六进制的字符串
3. 把十六进制的字符串 取出每两位字母的对应的十六进制的整形 与前两个字母对应的整形 进行异或运算
4. 返回最后一次异或计算结果
*/
func CheckMsg(s string) string {
	s = strings.ToUpper(hex.EncodeToString([]byte(s)))
	var temp int64 = 0
	var m int64 = 0
	lens := len(s)
	for i := 0; i < lens; i = i + 2 {
		m, _ = strconv.ParseInt(s[i:i+2], 16, 0)
		m = m ^ temp
		temp = m
	}
	return strconv.Itoa(int(temp))
}

/*
函数功能介绍：
1.
2.
3.
4.
*/
func InvokeWebservice(s string) string {

	return ""
}

// hash加密
package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// md5函数；参数：需要md5的字符串；返回值：md5后的字符串
func MD5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

// sha1函数；参数：需要sha1的字符串；返回值：sha1后的字符串
func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

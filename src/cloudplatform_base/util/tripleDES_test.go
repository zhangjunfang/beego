package util

import (
	//	"bytes"
	//	"crypto/cipher"
	//	"crypto/des"
	"encoding/base64"
	"fmt"
	"testing"
)

const ()

func init() {

}
func TestDesDecrypt(t *testing.T) {
	// DES 加解密
	testDes()
	// 3DES加解密
	test3Des()
}
func testDes() {
	key := []byte("sfe023f_")
	result, err := DesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

func test3Des() {
	key := []byte("sfe023f_sefiel#fi32lf3e!")
	result, err := TripleDesEncrypt([]byte("polaris@studygol"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := TripleDesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

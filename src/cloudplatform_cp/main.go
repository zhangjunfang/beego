// 1.实现云平台与慧修车服务端以及宇通之间通信协议 main.go
package main

import (
	"bufio"
	base "cloudplatform_base/base"
	"fmt"
	"net"
)

func main() {
	fmt.Println(base.LEFT_BRACKET)
	listenner, err := net.Listen("tcp", "127.0.0.1:9999")
	if nil != err {
		fmt.Printf("服务端获取监听程序出错：%v", err)
	} else {
		for {
			conn, err := listenner.Accept()
			fmt.Printf("%v , %s", conn, err)
			bufio.NewReader(conn)

			fmt.Printf("服务端收到： ")
		}
	}
}

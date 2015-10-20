// 1.实现云平台与慧修车服务端以及宇通之间通信协议 main.go
package main

import (
	"bufio"
	"fmt"
	//"math/rand"
	//base "cloudplatform_base"
	"io"
	"net"
	"strings"
)

func init() {

}

//const (
//	DOLLAR        string = "$"
//	LEFT_BRACKET  string = "["
//	RIGHT_BRACKET string = "]"
//	MSG_TYPE_R    string = "R"
//	MSG_TYPE_L    string = "L"
//	TYPE_L1       string = "L1"
//	//服务端应答类
//	TYPE_A     string = "$A$"
//	MSG_TYPE_A string = "A"
//)

func main() {
	listenner, err := net.Listen("tcp", "127.0.0.1:9999")
	defer listenner.Close()
	if nil != err {
		fmt.Printf("服务端获取监听-----程序出错：%v \r\n", err)
	} else {
		for {
			conn, err := listenner.Accept()
			if nil != err {
				fmt.Printf("服务端获取监听---接收程序出错：%v \r\n", err)
			}
			go HandleMessage(conn)
		}
	}

}

/*
 在模拟tcp粘包以及拆包过程完善后的代码
*/
func HandleMessage(conn net.Conn) {
	defer conn.Close()

	//注意这行代码，需要循环读取缓存中的数据
	reader := bufio.NewReader(conn) //原来在地61行 移动到这里 ---测试粘包过程中发现的问题
	for {
		s, err := reader.ReadString(']') //模拟拆包的过程中发现的问题---增加了for循环
		if nil != err {
			if _, ok := err.(*net.OpError); ok {
				fmt.Print("--------客户端断开连接------------------")
				return
			} else if io.EOF == err {
				fmt.Printf("服务端收到文件已经读到结尾： 是否需要释放资源")
				return
			} else {
				fmt.Printf("服务端收到错误信息： %T \r\n", err)
				fmt.Printf("服务端收到错误信息： %s \r\n", err)
			}
		}
		fmt.Println("收到客户端消息：：" + s)
		var msg []string = SplitString(s, DOLLAR, LEFT_BRACKET, RIGHT_BRACKET)
		handleMessage(conn, msg)
		//conn.Write([]byte(s))
		//fmt.Println("-----------------------------------------------------------------------------------------------")
	}

}
func handleMessage(conn net.Conn, s []string) {
	switch s[2] {
	//通用应答类消息
	case MSG_TYPE_R:
		{
			fmt.Println("通用应答类---------")
			conn.Write([]byte("通用应答类-----1----"))
		}
		//服务端应答类
	case MSG_TYPE_A:
		{
			//fmt.Println("服务端应答类---------")
			responseLoginMsg(conn, s)

		}
		//登录
	case MSG_TYPE_L:
		{
			if TYPE_L1 == s[3] {
				fmt.Println("登录---------")
				conn.Write([]byte("登录--L-L1------"))
			}

		}
	default:
		{
			fmt.Println("未知信息---------")
			conn.Write([]byte("未知信息---4------"))
		}
	}
}
func responseLoginMsg(conn net.Conn, s []string) {

	conn.Write([]byte("服务端应答类----2-----"))
}
func handleLand(conn net.Conn, s string) []string {

	conn.Write([]byte("已经登录平台了！！！！"))
	return nil

}

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

//开发使用的测试例子 ， 生产环境不建议使用
func Test(s string) []string {
	//var s string = "[ef7855239d1547339492a898a746ed34$1$R$$1440749683041$1$30]"
	s = strings.TrimLeft(s, LEFT_BRACKET)
	fmt.Printf(strings.TrimRight(s, RIGHT_BRACKET))
	//基于性能考虑 暂时不适用上面的方式
	var arrays []string = strings.Split(s, DOLLAR)
	//下面同样存在这样的问题
	arrays[0] = strings.TrimLeftFunc(arrays[0], func(c rune) bool {
		if LEFT_BRACKET == string(c) {
			return true
		}
		return false
	})
	var lens int = len(arrays) - 1
	arrays[lens] = strings.TrimRightFunc(arrays[lens], func(c rune) bool {
		if RIGHT_BRACKET == string(c) {
			return true
		}
		return false
	})
	return arrays
}

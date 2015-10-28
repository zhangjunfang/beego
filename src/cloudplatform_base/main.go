/*
1.实现云平台与慧修车服务端以及宇通之间通信协议 main.go
2.golang中导入包 是相对src目录的相对路径和实际的包名没有直接的关系
*/
package main

import (
	"bufio"
	"fmt"
	//"math/rand"
	"bytes"
	base "cloudplatform_base/base"
	util "cloudplatform_base/util"
	"io"
	"net"
	"runtime"
	"strings"
	"time"
)

func init() {
	//runtime.GOMAXPROCS(256)
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.GOOS, runtime.GOROOT(), runtime.GOARCH, runtime.Version())
}

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
		var msg []string = util.SplitString(s, base.DOLLAR, base.LEFT_BRACKET, base.RIGHT_BRACKET)
		handleMessage(conn, msg)
		//conn.Write([]byte(s))
		//fmt.Println("-----------------------------------------------------------------------------------------------")
	}

}
func handleMessage(conn net.Conn, s []string) {
	switch s[2] {
	//通用应答类消息
	case base.MSG_TYPE_R:
		{
			fmt.Println("通用应答类---------")
			conn.Write([]byte("通用应答类-----1----"))
		}
		//服务端应答类
	case base.MSG_TYPE_A:
		{
			//fmt.Println("服务端应答类---------")
			responseLoginMsg(conn, s)

		}
		//登录
	case base.MSG_TYPE_L:
		{
			if base.TYPE_L1 == s[3] {
				fmt.Println("登录---------")
				conn.Write([]byte("登录--L-L1------"))
			}

		}
	default:
		{
			fmt.Println("未知信息---------")
			conn.Write([]byte("未知信息[数据类型]---4------"))
		}
	}
}

/*
1.测试数据：[ef7855239d1547339492a898a746ed34$1$A$$1440749683041$1$30]
*/
func responseLoginMsg(conn net.Conn, s []string) {
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	buffer.WriteString(base.LEFT_BRACKET)
	buffer.WriteString(s[0])
	buffer.WriteString(base.DOLLAR)
	buffer.WriteString(s[1])
	//			if (arr[3].equals(Constant.TYPE_L1)) { // 登录验证
	//				buff.append(getLoginResponse(session,arr) + Constant.DOLLAR);
	//			}
	if base.MSG_TYPE_L == s[2] {
		buffer.WriteString(base.TYPE_A)
		buffer.WriteString(base.DOLLAR)
		buffer.WriteString(strings.Replace(s[3], "L", "A", -1))
		buffer.WriteString(base.DOLLAR)
		buffer.WriteString(time.Now().String())
		buffer.WriteString(base.DOLLAR)
		if s[3] == base.TYPE_L1 { // 登录验证
			buffer.WriteString(time.Now().String())
			buffer.WriteString(base.DOLLAR)
		} else if s[3] == base.TYPE_L2 {
			buffer.WriteString(base.N1)
			buffer.WriteString(base.DOLLAR)
		}
	}
	buffer.WriteString(base.DOLLAR)
	temp := buffer.String()
	buffer.WriteString(util.CheckMsg(temp[1 : len(temp)-1])) // 校验码
	buffer.WriteString(base.RIGHT_BRACKET)
	fmt.Println("拼接后的结果为-->", buffer.String())
	conn.Write([]byte(buffer.String()))

	//	conn.Write([]byte("服务端应答类----3-----"))
}
func handleLand(conn net.Conn, s string) []string {

	conn.Write([]byte("已经登录平台了！！！！"))
	return nil

}

//开发使用的测试例子 ， 生产环境不建议使用
func Test(s string) []string {
	//var s string = "[ef7855239d1547339492a898a746ed34$1$R$$1440749683041$1$30]"
	s = strings.TrimLeft(s, base.LEFT_BRACKET)
	fmt.Printf(strings.TrimRight(s, base.RIGHT_BRACKET))
	//基于性能考虑 暂时不适用上面的方式
	var arrays []string = strings.Split(s, base.DOLLAR)
	//下面同样存在这样的问题
	arrays[0] = strings.TrimLeftFunc(arrays[0], func(c rune) bool {
		if base.LEFT_BRACKET == string(c) {
			return true
		}
		return false
	})
	var lens int = len(arrays) - 1
	arrays[lens] = strings.TrimRightFunc(arrays[lens], func(c rune) bool {
		if base.RIGHT_BRACKET == string(c) {
			return true
		}
		return false
	})
	return arrays
}

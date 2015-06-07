// learn
package learn

import (
	"fmt"
	"runtime"
	//"strings"
)

/**
*全局变量【包含常量】定义 但不一定使用  符合golang语法
* 局部变量 定义必须使用
*
 */
var (
	a int
	b int
	d string
)

/*
*多变量赋值： 首先计算相关之的值，从左到右依次赋值
* 常量: 必须是编译期 确定的值  数字  字符串  布尔值 len  cap  unsafe.Sizeof
 */
var (
	q int32
	m string
	r []int
)
/*
  字符串的操作说明：
     1.子字符串指向原来的字节数组
     2.修改了指针和长度属性
  字符串的转换需要重新分配内存地址以及复制字节数组
**/
func Test() {
	fmt.Println("------------------------------3----------------------------!")
	var m string = "张伯雨"
	fmt.Println(m[:])
	fmt.Println(m[0:3])
	fmt.Println(m[3:6])
	//输出变量m的数据类型
	fmt.Printf("%T\r\n", m)
	//输出字符串内存地址
	fmt.Println("\r\n", &m)
/*
golang 中的for 循环分为两种情况：
   1.基于字节遍历 使用字符串的len方法 使用的是字节遍历
   2.基于字符遍历 使用字符串的range 使用的是字符遍历
*/
	for i:=0;i<len(m);i++{
		fmt.Printf("%d , %c \r\n",i,m[i])
	 }
	fmt.Println("------------------------------33----------------------------!")
	for  index , v :=range m{
		fmt.Printf("%d , %c \r\n",index,v)
	}
}
func TestVar() {
	fmt.Println("------------------------------2----------------------------!")
	d := "fdfdf"
	//使用语言自带的内置方法
	fmt.Println(len(d))
	fmt.Println(d)
	v := make([]int, 2, 8)
	for index, v := range v {
		fmt.Println(index, v)
	}
	//基本数据类型之间相互转换
	fmt.Println(string([]byte(d)))
}

func TestRuntime() {
	fmt.Println("------------------------------1----------------------------!")
	fmt.Println("Hello learnner!")
	//runtime.Gosched
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	runtime.LockOSThread()
	fmt.Println(runtime.MemProfileRate)
	fmt.Println(runtime.NumCgoCall())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.Version())
}

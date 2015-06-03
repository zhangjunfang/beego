// learn
package learn

import (
	"fmt"
	"runtime"
	//"strings"
)
/**
*全局变量定义 但不一定使用
* 局部变量 定义必须使用
*
*/
var (
	a  int  
	b  int  
	d  string
	
)
/*
*多变量赋值： 首先计算相关之的值，从左到右依次赋值
* 常量 必须是编译期 确定的值  数字  字符串  布尔值
*  len  cap  unsafe.Sizeof
*/
var (
	q int32
	m string
	r  []int
)
func Test(){
	
}
func  TestVar(){
    fmt.Println("------------------------------2----------------------------!")
	d:="fdfdf"
	//使用语言自带的内置方法
	fmt.Println(len(d))
	fmt.Println(d)
	v:=make([]int, 2,8)
	for index , v := range v{
		fmt.Println(index , v )
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


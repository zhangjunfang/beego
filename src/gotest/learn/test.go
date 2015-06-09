// learn
package learn

import (
	"fmt"
	"runtime"
	//"strings"
)

func Test() {
	/*
		闭包复制的是原对象的指针，这就容易解释延迟引用对象
		当且仅当 最后一个panic 错误 被捕获
	*/
	/**
	  type：
	  可以在全局或者局部函数中定义新的类型.
	  新类型不是原来类型的别名，
	  除了拥有相同的数据存储结构外，他们之间没有必然的联系
	*/
	fmt.Println("------------------------------4----------------------------!")

	type bite int
	var x bite = 100
	fmt.Sprintf("%s", x)
	fmt.Println(x)
	/*
				golang中保留关键之字如下：
				break  case  chan const continue
				default defer else fallthrough  for
		        func go  goto if  import
				interface map package range return
				select struct  switch type var
				****/
	//初始化语句未必是定义变量

	a := []int{1, 2, 3}
	for index, value := range a {
		if index == 0 {
			a[0] = 111
			a[1] = 222
			fmt.Println(a)
			fmt.Println(value)
		}
		a[index] = value + 400
		fmt.Println(&a)
	}
	fmt.Println(&a)
	fmt.Printf("%T\r\n", a)
	fmt.Println(a)
}

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
*多变量赋值： 首先计算相关的值，从左到右依次赋值
* 常量: 必须是编译期确定的值  数字  字符串  布尔值 len  cap  unsafe.Sizeof
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
func TestCast() {
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
	for i := 0; i < len(m); i++ {
		fmt.Printf("%d , %c \r\n", i, m[i])
	}
	fmt.Println("------------------------------33----------------------------!")
	for index, v := range m {
		fmt.Printf("%d , %c \r\n", index, v)
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

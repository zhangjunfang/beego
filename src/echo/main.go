package main

import (
	"unsafe"
)

const (
	a = "abc"

	b = len(a)

	c = unsafe.Sizeof(b)
)
const (
	s = "abc"

	x // x = "abc"

)

func main() {
	println(x)
	const xx = "fdsdfsdfs"
	s := "abc"

	println(&s)

	s, y := "hello", 20 // 重新赋值: 与前 s 在同一层次的代码块中，且有新的变量被定义。

	println(&s, y) // 通常函数多返回值 err 会被重复使?。

	{

		s, z := 1000, 30 // 定义新同名变量: 不在同一层次代码块。

		println(&s, z)

	}
}

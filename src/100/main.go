package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
)

func NetTest() {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, intes := range interfaces {
			fmt.Println("intes fsdfsdf ", intes.HardwareAddr.String())
			//if intes == nil {
			//	fmt.Println("sfdfsf  nil")
			//} else {
			//	fmt.Println("sfdfsf" + intes.Name)
			//}

		}
	}

}

func main() {
	fmt.Println("Hello World!")
	NetTest()
	fmt.Println("====================os=================================")
	fmt.Println(os.Hostname())
	p, _ := os.FindProcess(1260)
	fmt.Println(strconv.Itoa(p.Pid))

	fmt.Println(strconv.Itoa(p.Pid))
	fmt.Println("====================runtime=================================")
	//pc uintptr, file string, line int, ok bool
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc, file, line, ok)
	fmt.Printf(string(runtime.GOMAXPROCS(runtime.NumGoroutine())))
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumGoroutine())
	fun := runtime.FuncForPC(pc)
	fmt.Println(fun.Name())
	fmt.Println(fun.Entry())
	fmt.Println(runtime.GOROOT())
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m) //获取系统资源
	fmt.Println(m.TotalAlloc)
}

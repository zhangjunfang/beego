// nettest project main.go
package main

import (
	"fmt"
	"net"
	"os"
)

func NetTest() {
	interfaces, err := net.Interfaces()
	fmt.Println(net.InterfaceByIndex(2))
	if err == nil {
		for _, intes := range interfaces {
			addrs, _ := intes.Addrs()
			for _, add := range addrs {
				fmt.Println("network:", add.Network())
			}
			fmt.Println("HardwareAddr:", intes.HardwareAddr.String())
			fmt.Println("name:" + intes.Name)

		}
	}

}
func Ostest() {
	os.NewFile()
}
func main() {
	fmt.Println("Hello World!")
	//NetTest()
}

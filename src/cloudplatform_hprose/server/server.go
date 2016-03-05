package server

import (
	"errors"
	"fmt"
	"github.com/hprose/hprose-go"
	"net/http"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

type myService struct{}

func (myService) Swap(a int, b int) (int, int) {
	return b, a
}

func (myService) Sum(args ...int) (int, error) {
	if len(args) < 2 {
		return 0, errors.New("Requires at least two parameters")
	}
	a := args[0]
	for i := 1; i < len(args); i++ {
		a += args[i]
	}
	return a, nil
}

func SynchronousInvokingServer() {
	fmt.Println("--------------------------")
	service := hprose.NewHttpService()
	service.AddFunction("hello", hello)
	service.AddMethods(myService{})
	http.ListenAndServe(":8080", service)

}

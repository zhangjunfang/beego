package main

import (
	"api/controllers"
	_ "api/routers"
	"github.com/astaxie/beego"
	//"sync"
)

func main() {

	beego.Router("/page", &controllers.MainController{}, "get,post:GetMobileJson")
	beego.Router("/ip", &controllers.MainController{}, "get,post:GetIPDesc")
	beego.Router("/mobile/:phone([^/]+)", &controllers.MainController{}, "get,post:GetMobileJson")
	//sync.WaitGroup{}
	beego.Run()
}

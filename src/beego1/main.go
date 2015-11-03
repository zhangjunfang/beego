package main

import (
	_ "beego1/routers"
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/idada/v8.go"
	//_ "github.com/bitly/go-simplejson"
	//_ "github.com/realint/dbgutil"
)

//参见config.go查询
//默认 充分利用cpu的资源
func main() {
	//session 配置  默认存在内存中
	beego.SessionOn = true
	beego.SessionName = "jsessionId"
	beego.SessionProvider = "redis"
	beego.SessionSavePath = "127.0.0.1:6379"
	//日志设置
	beego.SetLevel(0)
	beego.SetLogFuncCall(true)
	//在页面中 显示错误
	beego.ErrorsShow = true
	//xss配置
	beego.EnableXSRF = true
	beego.XSRFKEY = "多个地方个地方广泛的个地方广泛的个地方"
	beego.XSRFExpire = 3600 * 5 //5m
	// 显示服务器 名称
	beego.BeegoServerName = "ocean-zby"
	//展示服务器API或者docs
	beego.EnableDocs = true
	//获取当前项目所在的物理路径
	fmt.Println("AppPath:::", beego.AppPath)
	//程序的入口
	beego.Run("127.0.0.1:9090")

}

package main

import (
	_ "bootstrap/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func main() {
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/img", "static/img")
	//注册国际化配置信息
	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	//日志级别
	beego.SetLevel(0)
	//注册模板函数【国际化】
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}

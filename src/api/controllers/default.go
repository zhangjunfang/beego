package controllers

import (
	"api/models"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *MainController) GetMobileJson() {
	//phone := this.Ctx.Input.Param(":phone")
	phone := this.GetString("phone")
	fmt.Println("phone:", phone)
	json := models.GetMobileDesc(phone)
	this.Data["json"] = json
	this.ServeJson()
}

func (this *MainController) GetIPDesc() {
	ip := this.GetString("ip")
	ips := models.GetIPDesc(ip)
	fmt.Println(ips.Root)
	for k, v := range ips.Eelement {
		fmt.Println(k, "::", v.Root)
	}
	this.TplNames = "index.tpl"
}

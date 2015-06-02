package routers

import (
	"beego1/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/adduser", &controllers.MainController{}, "get,post:AddUser")
	beego.Router("/login", &controllers.MainController{}, "get,post:Login")
	beego.Router("/login2", &controllers.MainController{}, "get,post:Post")
	beego.Router("/redirect", &controllers.MainController{}, "get,post:Forward")
	beego.Router("/update", &controllers.MainController{}, "get,post:Update")
	beego.Router("/updatebatch", &controllers.MainController{}, "get,post:UpdateBatch")
	beego.Router("/deletebatch", &controllers.MainController{}, "get,post:DeleteBatchPerson")
	beego.Router("/delete", &controllers.MainController{}, "get,post:DeletePerson")
	beego.Router("/upload", &controllers.MainController{}, "get,post:FileUpload")
	beego.Router("/httpredirect", &controllers.MainController{}, "get,post:HttpRedirect")
	beego.Router("/httpdown", &controllers.MainController{}, "get,post:FileDown")
	beego.Router("/sessiontest", &controllers.MainController{}, "get,post:SessionTest")
	beego.Router("/json", &controllers.MainController{}, "get,post:Tojson")
	beego.Router("/xml", &controllers.MainController{}, "get,post:Toxml")
	beego.Router("/jsonp", &controllers.MainController{}, "get,post:Tojsonp")
	beego.Router("/page", &controllers.MainController{}, "get,post:Topage")
}

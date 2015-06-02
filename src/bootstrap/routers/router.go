package routers

import (
	"bootstrap/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//国际化例子
	beego.Router("/", &controllers.MainController{})
	//获取服务器这个系统目录结构【每次加载一个目录下的文件】
	beego.Router("/fileList", &controllers.MainController{}, "get,post:FileList")
	//获取服务器全部的目录结构【一次性加载】
	beego.Router("/fileAllList", &controllers.MainController{}, "get,post:FileAllList")
	//远程执行服务端命令【参数和命名一次性传入】
	beego.Router("/executeCommand", &controllers.MainController{}, "get,post:ExecuteCommand")
	//远程执行服务端命令【参数和命名之间使用英文分号分割】
	beego.Router("/executeCommandWithArgs", &controllers.MainController{}, "get,post:ExecuteCommandWithArgs")
	//文件上传
	beego.Router("/fileUpload", &controllers.MainController{}, "get,post:FileUpload")
	//文件下载
	beego.Router("/fileDown", &controllers.MainController{}, "get,post:FileDown")
	//文件移动
	beego.Router("/moveFile", &controllers.MainController{}, "get,post:MoveFile")
	//SendMC
	beego.Router("/sendMC", &controllers.MainController{}, "get,post:SendMC")
}

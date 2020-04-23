package routers

import (
	"hello2/controllers"

	"github.com/astaxie/beego"
)

//init方法设置路由信息
func init() {
	//如果URL是/，则调用后面的Controler
	beego.Router("/", &controllers.MainController{}) //默认调用get方法
	//beego.Router("/hi", &controllers.MainController{}, "get:Hi")
	//beego.Router("/user",&controllers.UserController{},"get:GetInfo")
	beego.Router("/user/?:id",&controllers.UserController{},"get:GetInfo")
	beego.Router("/user", &controllers.UserController{},"get:GetFile;post:PostFile")
		//beego.Router("/user",&controllers.UserController{})
	//hi请求来了会调用HI方法
}


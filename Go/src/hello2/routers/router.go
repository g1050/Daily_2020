package routers

import (
	"hello2/controllers"

	"github.com/astaxie/beego"
)

//init方法设置路由信息
func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hi", &controllers.MainController{}, "get:Hi")
	//hi请求来了会调用HI方法
}

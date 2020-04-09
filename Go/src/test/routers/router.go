package routers

import (
	"test/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hi", &controllers.MainController{}, "get:Hi")

}

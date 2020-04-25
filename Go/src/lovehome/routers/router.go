package routers

import (
	"lovehome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{},"Get:GetArea")
	beego.Router("/api/v1.0/session", &controllers.SessionController{},"Get:MyGetSession")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"Get:GetHouseIndex")
	beego.Router("/api/v1.0/users", &controllers.UserController{},"Post:PostRegisterData")

}

package routers

import (
	"lovehome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{},"Get:GetArea")
	beego.Router("/api/v1.0/session", &controllers.SessionController{},"Get:MyGetSession;Delete:MyDeleteSession")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"Get:GetHouseIndex")
	beego.Router("/api/v1.0/users", &controllers.UserController{},"Post:PostRegisterData")
	beego.Router("/api/v1.0/sessions", &controllers.UserController{},"Post:PostLoginData")
	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{},"Post:PostAvatar")
	beego.Router("/api/v1.0/user", &controllers.UserController{},"Get:GetUserData")
	beego.Router("/api/v1.0/user/name", &controllers.UserController{},"Put:PutNameData")
	beego.Router("/api/v1.0/user/auth", &controllers.UserController{},"")
}

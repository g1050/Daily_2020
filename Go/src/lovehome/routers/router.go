package routers

import (
	"fmt"
	"lovehome/controllers"
	"github.com/astaxie/beego"
)

func init() {

	fmt.Println("router ----------")
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{},"Get:GetArea")
	beego.Router("/api/v1.0/session", &controllers.SessionController{},"Get:MyGetSession;Delete:MyDeleteSession")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseController{},"Get:GetHouseIndex")
	beego.Router("/api/v1.0/users", &controllers.UserController{},"Post:PostRegisterData")
	beego.Router("/api/v1.0/sessions", &controllers.UserController{},"Post:PostLoginData")
	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{},"Post:PostAvatar")
	beego.Router("/api/v1.0/user", &controllers.UserController{},"Get:GetUserData")
	beego.Router("/api/v1.0/user/name", &controllers.UserController{},"Put:PutNameData")
	beego.Router("/api/v1.0/user/auth", &controllers.UserController{},"Get:GetUserData;Post:PostRealName")
	beego.Router("/api/v1.0/user/houses", &controllers.HouseController{},"Get:GetMyHouseData")
	beego.Router("/api/v1.0/houses", &controllers.HouseController{},"Post:PostHouseData")
	beego.Router("/api/v1.0/houses/*/images", &controllers.HouseController{},"Post:PostHouseImage")
	beego.Router("/api/v1.0/houses/?:house_id", &controllers.HouseController{},"Get:GetHouseDetail")
}

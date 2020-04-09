package controllers

import (
	"fmt"
	"test/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//自定义方法,用户有请求会调用该方法
func (c *MainController) Hi() {
	fmt.Println("Hello gxk")
	models.PrintUsers()
	c.TplName = "index.tpl"
}

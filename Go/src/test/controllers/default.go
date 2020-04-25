package controllers

import (
	"fmt"
	//"test/models"
	"astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("Hello world")
}

//自定义方法,用户有请求会调用该方法
func (c *MainController) Hi() {
	fmt.Println("Hello gxk")

	//存储session的key和value
	c.SetSession("cmsusername", "Jack")
	user := c.GetSession("cmsusername")
	logs.Informational("username:Jack loged in")
	fmt.Println("user----->", user) //session 的使用

	//models.PrintUsers()
	c.TplName = "index.tpl"
}

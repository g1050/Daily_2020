package controllers

import (
	"astaxie/beego"
	"fmt"
	"hello2/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Hi() {
	fmt.Println("Hi console")
	//models.PrintUsers()
	models.PrintUserByORM()
	c.TplName = "index.tpl"
}

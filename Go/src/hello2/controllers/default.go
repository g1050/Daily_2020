package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"hello2/models"
)

type MainController struct {
	beego.Controller //包含关系，类似继承
}



//重写Get方法
func (c *MainController) Get() {
	//models.InsertData()
	//models.DeleteData()
	//models.UpdateDate()
	//models.Read()
	//models.InsertOrder()
	c.Data["Website"] = "baidu.com"
	c.Data["Email"] = "1005035266@qq.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Hi() {
	fmt.Println("Hi console")
	//models.PrintUsers()
	models.PrintUserByORM()
	c.TplName = "index.tpl"
}

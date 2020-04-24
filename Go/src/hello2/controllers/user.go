package controllers
import (
	"astaxie/beego"
)

type UserController struct {
	beego.Controller //包含关系，类似继承
}

func (c *UserController) Get() {
	//c.Data["Website"] = "baidu.com"
	//c.Data["Email"] = "1005035266@qq.com"
	//c.TplName = "index.tpl"
	c.Ctx.WriteString("Get suvessfuly!")
}

func (c *UserController) GetFile(){
	path := c.Ctx.Input.Param(":path")
	ext := c.Ctx.Input.Param(":ext")

	c.Ctx.WriteString("Get data path = "+path + " ext = "+ext)
}
func (c *UserController) GetInfo() {
	c.Ctx.WriteString("Hello World\n")
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("Get data id = "+id)
	//fmt.Println("Hi console")
	//models.PrintUsers()
	//models.PrintUserByORM()
	//c.TplName = "index.tpl"
}

func (c *UserController) PostFile() {
	c.Ctx.WriteString("This is post function")
}


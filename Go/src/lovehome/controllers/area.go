package controllers

import (
	"astaxie/beego"
	"astaxie/beego/logs"
	"lovehome/models"
	_"lovehome/models"
)
type AreaController struct {
	beego.Controller
}

func (c *AreaController) sendJSON(mp map[string]interface{}){
	c.Data["json"]=mp
	c.ServeJSON()
}

func (c *AreaController) GetArea() {
	//从数据库中获取数据
	//c.Ctx.WriteString("Hello")
	mp := models.SelectAreaDate()
	logs.Info(mp)
	//打包成json格式发给客户端
	c.sendJSON(mp)
}
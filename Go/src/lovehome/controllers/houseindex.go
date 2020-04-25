package controllers

import (
	"astaxie/beego"
	"astaxie/beego/logs"
)

type HouseIndexController struct {
	beego.Controller
}

func (c *HouseIndexController) sendJSON(mp map[string]interface{}){
	c.Data["json"]=mp
	c.ServeJSON()
}

func (c *HouseIndexController) GetHouseIndex() {
	//从数据库中获取数据
	//c.Ctx.WriteString("Hello")
	//mp := models.SelectAreaDate()
	mp := make(map[string]interface{})
	mp["errno"]	 = 404
	logs.Info(mp)
	//打包成json格式发给客户端
	c.sendJSON(mp)
}

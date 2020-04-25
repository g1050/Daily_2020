package controllers

import (
	"astaxie/beego"
	"astaxie/beego/logs"
	"lovehome/models"
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
	ok,areas := models.SelectAreaDate()

	mp := make(map[string]interface{})
	//打包成json格式发给客户端
	defer c.sendJSON(mp)

	//将返回的area打包成map
	if ok{
		mp["errno"] = models.RECODE_OK
		mp["errmsg"] = models.RecodeText(models.RECODE_OK)
		mp["data"] = areas
	}else{
		logs.Error("未读取到Area数据")
		mp["errno"] = models.RECODE_DBERR
		mp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	}
}
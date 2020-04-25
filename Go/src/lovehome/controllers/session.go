package controllers

import (
	"astaxie/beego"
	"lovehome/models"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) sendJSON(mp map[string]interface{}){
	c.Data["json"]=mp
	c.ServeJSON()
}

func (c *SessionController) MyGetSession() {
	//从数据库中获取数据
	//c.Ctx.WriteString("Hello")
	//mp := models.SelectAreaDate()
	user := models.User{}

	mp := make(map[string]interface{})
	//打包成json格式发给客户端
	defer c.sendJSON(mp)

	//设置错误信息
	mp["errno"] = models.RECODE_DBERR
	mp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	//获取session
	name := c.GetSession("name")
	if name != nil{
		user.Name = name.(string)
		mp["errno"] = models.RECODE_OK
		mp["errmsg"] = models.RecodeText(models.RECODE_OK)
		mp["data"] = user
	}

	//logs.Info(mp)
}

func (c *SessionController) MyDeleteSession(){
	mp := make(map[string]interface{})
	defer c.sendJSON(mp)

	c.DelSession("name")
	mp["errno"] = models.RECODE_OK
	mp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

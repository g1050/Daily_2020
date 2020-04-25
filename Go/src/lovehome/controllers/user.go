package controllers

import (
	"astaxie/beego"
	"astaxie/beego/logs"
	"encoding/json"
	"lovehome/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) sendJSON(mp map[string]interface{}){
	c.Data["json"]=mp
	c.ServeJSON()
}

func (c *UserController) PostRegisterData() {
	mp := make(map[string]interface{})
	//获取前端的数据
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&mp);err != nil{
		logs.Info(err)
	}

	//设置session
	c.SetSession("name",mp["mobile"].(string))
	//插入数据库中,这里的mp用来存储发来的数据

	models.InsertUserData(mp)
	//logs.Info(ok)
	ok := true;

	mpreturn := make(map[string]interface{})
	defer c.sendJSON(mpreturn)
	//defer logs.Info(mpreturn)
	//后面的mp用来存储发送过去的数据
	if ok == false{
		mpreturn["errno"] = "4002"
		mpreturn["errmsg"] = "注册失败"
	}else{
		mpreturn["errno"] = "0"
		mpreturn["errmsg"] = "注册成功"
	}
}

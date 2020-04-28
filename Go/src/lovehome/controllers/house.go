package controllers

import (
	"astaxie/beego"
	"astaxie/beego/logs"
	"fmt"
	"lovehome/models"
)

type HouseController struct {
	beego.Controller
}

func (c *HouseController) sendJSON(mp map[string]interface{}){
	c.Data["json"]=mp
	c.ServeJSON()
}

func (c *HouseController) GetMyHouseData(){
	//返回数据
	mp := make(map[string]interface{})
	defer c.sendJSON(mp)

	//从session中获取id
	id := c.GetSession("id")
	//从数据库中取出对应useid的数据
	ok,houses  := models.SelectHouseDataByUserID(id.(int))
	if !ok {
		mp["errno"] = models.RECODE_DBERR
		mp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//house转map
	housemap := make(map[string]interface{})
	housemap["house"] = houses
	mp["errno"] = models.RECODE_OK
	mp["errmsg"] = models.RecodeText(models.RECODE_OK)
	mp["data"] = housemap
	fmt.Println(houses)
}

func (c *HouseController) GetHouseIndex() {
	//从数据库中获取数据
	//c.Ctx.WriteString("Hello")
	//mp := models.SelectAreaDate()
	mp := make(map[string]interface{})
	mp["errno"]	 = 404
	logs.Info(mp)
	//打包成json格式发给客户端
	c.sendJSON(mp)
}

package controllers

import (
	"astaxie/beego"
	"astaxie/beego/logs"
	"encoding/json"
	"fmt"
	"lovehome/models"
	"strconv"
)

type HouseController struct {
	beego.Controller
}

func (c *HouseController) sendJSON(mp map[string]interface{}){
	c.Data["json"]=mp
	c.ServeJSON()
}

func (c *HouseController)PostHouseImage()  {
	resp := make(map[string]interface{})
	defer c.sendJSON(resp)
	datamap:= make(map[string]string)

	//获取session中的houseid
	houseid := c.GetSession("houseid").(int)

	//获得图片数据,给M存储并且返回URL

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	datamap["url"] = "http://192.168.0.115:8080/group1/default/20200427/19/17/5/Hummingbird_by_Shu_Le.jpg"
	resp["data"] = datamap

}

func (c *HouseController) PostHouseData(){
	resp := make(map[string]interface{})
	defer c.sendJSON(resp)

	//从前端获取数据
	datamap := make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody,&datamap)
	//判断前段数据的合法性,其实所有请求都要检查
	fmt.Println("datamap",datamap)

	//插入数据到数据库
	house := models.House{}
	fmt.Println("dot.....................................1")
	//Area
	tmp,_ := strconv.Atoi(datamap["area_id"].(string))
	area := models.Area{Id: tmp,Name: "武安"}
	house.Area = &area
	//User
	user := models.User{Id: c.GetSession("id").(int)}
	house.User = &user
	//Facilities复杂操作
	facilities := []models.Facility{}
	for _,f := range datamap["facility"].([]interface{}) {
		//获得当个设施ID
		fid,_ := strconv.Atoi(f.(string))
		//创建单个设施
		fac := models.Facility{Id: fid}
		//添加进数组
		facilities = append(facilities,fac)
	}

	//othres
	//string转int,需要辅助数字
	tmp,_ = strconv.Atoi(datamap["acreage"].(string))
	house.Acreage = tmp
	house.Address = datamap["address"].(string)
	house.Beds = datamap["beds"].(string)
	tmp,_ = strconv.Atoi(datamap["capacity"].(string))
	house.Capacity = tmp
	//Ctime?
	tmp,_ = strconv.Atoi(datamap["deposit"].(string))
	house.Deposit = tmp
	tmp,_ = strconv.Atoi(datamap["price"].(string))
	house.Price = tmp
	tmp,_ = strconv.Atoi(datamap["room_count"].(string))
	house.Room_count = tmp
	house.Title = datamap["title"].(string)
	house.Unit = datamap["unit"].(string)
	tmp,_ = strconv.Atoi(datamap["max_days"].(string))
	house.Max_days = tmp
	tmp,_ = strconv.Atoi(datamap["min_days"].(string))
	house.Min_days = tmp


	//插入house和facility
	ok,id := models.InsertHouseData(house,facilities)
	if !ok {
		resp["errno"] = models.RECODE_DBERR
		resp["errno"] = models.RecodeText(models.RECODE_DBERR)
	}
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	mp:= make(map[string]interface{})
	mp["house_id"]	= id
	//设置houseid的session
	c.SetSession("houseid",id)
	resp["data"] = mp

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
	housemap["houses"] = houses
	mp["errno"] = models.RECODE_OK
	mp["errmsg"] = models.RecodeText(models.RECODE_OK)
	mp["data"] = housemap
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

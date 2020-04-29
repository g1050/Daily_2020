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

func (c *HouseController) GetHouseDetail(){
	resp := make(map[string]interface{})
	defer c.sendJSON(resp)

	//获取user_id和house_id
	tmp := c.Ctx.Input.Param(":house_id")
	house_id,_ := strconv.Atoi(tmp)
	user_id := c.GetSession("id").(int)
	logs.Info("user_id = ",user_id,"house_id",house_id)

	//从缓存中取出房屋数据

	//关联查询
	//house获取
	housemap := make(map[string]interface{})
	house := models.SelectHouseDataByHouseId(house_id)
	_,user := models.SelectUserDataById(user_id)
	house.User = &user

	/*
	resp["facilities"].appen


	facilities := []models.Facility{}
	for _,f := range datamap["facility"].([]interface{}) {
		//获得当个设施ID
		fid,_ := strconv.Atoi(f.(string))
		//创建单个设施
		fac := models.Facility{Id: fid}
		//添加进数组
		facilities = append(facilities,fac)
	}
	 */
	housemap["house"] = house
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = housemap
}

func (c *HouseController)PostHouseImage()  {
	resp := make(map[string]interface{})
	defer c.sendJSON(resp)
	datamap:= make(map[string]string)

	//获得图片数据,给M存储并且返回URL
	filedata,hd,err := c.GetFile("house_image");
	hid,_ := c.GetInt("house_id")
	fmt.Println("house_id ===== ",hid)
	if err != nil{
		resp["errno"] =  models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		logs.Error("获取二进制数据失败")
		return
	}

	//发给M层存储到fastdfs中
	bufer := make([]byte,hd.Size)
	filedata.Read(bufer)
	path := models.UploadFile(bufer,hd.Filename)

	//把url打包发给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	datamap["url"] = path
	resp["data"] = datamap

	//根据hid获取house数据
	house := models.SelectHouseDataByHouseId(hid)
	//向houseimage中插入数据
	houseimage := models.HouseImage{}
	houseimage.Url = path
	houseimage.House = &house
	models.InsertHouseImage(houseimage)
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

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

func (c *UserController) PostRealName(){
	//返回数据给前端
	resp := make(map[string]interface{})
	defer c.sendJSON(resp)

	//获取前端数据
	datamap:= make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody,&datamap)
	//得到session中的id
	id := c.GetSession("id")
	//更新数据库中的user数据
	user := models.User{Id: id.(int),Real_name: datamap["real_name"].(string),Id_card: datamap["id_card"].(string)}
	err := models.UpdateUserData(user,"id_card")
	err = models.UpdateUserData(user,"real_name")
	if err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		logs.Error(err)
	}
	//构造resp
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = datamap
}


func (c *UserController) PutNameData(){
	//返回数据给前端
	resp := make(map[string]interface{})
	defer c.sendJSON(resp)

	//获取前端数据
	datamap:= make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody,&datamap)
	//得到session中的id
	id := c.GetSession("id")
	//更新数据库中的user数据
	user := models.User{Id: id.(int),Name: datamap["name"].(string)}
	err := models.UpdateUserData(user,"name")
	//更新session中的name
	if err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		logs.Error(err)
	}
	c.SetSession("name",datamap["name"].(string))
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = datamap
	//构造resp

}
func (c *UserController) GetUserData(){

	resp := make(map[string]interface{})
	//把数据打包,发送
	defer c.sendJSON(resp)

	//获取session
	id := c.GetSession("id")
	//根据session里的id得到数据
	ok,user := models.SelectUserDataById(id.(int))
	if ok{
		//给resp赋值
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"]  = user
	}else{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	}

}

func (c *UserController) PostAvatar(){

	//获得文件的二进制数据data,文件信息hd
	mp := make(map[string]interface{})
	defer c.sendJSON(mp)
	//前段发来的二进制数据的字段就叫avatar
	filedata,hd,err := c.GetFile("avatar");
	if err != nil{
		mp["errno"] =  models.RECODE_REQERR
		mp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		logs.Error("获取二进制数据失败")
		return
	}

	//发给M层存储到fastdfs中
	bufer := make([]byte,hd.Size)
	filedata.Read(bufer)
	path := models.UploadFile(bufer,hd.Filename)

	//把url打包发给前端
	mp["errno"] = models.RECODE_OK
	mp["errmsg"] = models.RecodeText(models.RECODE_OK)
	url := make(map[string]string)
	url["avatar_url"] = path
	mp["data"] = url

	//更新user中的数据
	userUp := models.User{Id: c.GetSession("id").(int),Avatar_url: path}
	models.UpdateUserData(userUp,"avatar_url")
	//获取session中的

	/*
	//获取后缀名字
	suffix := path.Ext(hd.Filename)
	logs.Info("suffix",suffix)
	//存储文件到fastdfs中
	 */




}

func (c *UserController)PostLoginData() {
	mp := make(map[string]interface{})
	defer c.sendJSON(mp)

	//获取前段发送的账号信息
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&mp);err != nil {
		logs.Info(err)
	}
	//获取数据库中的信息
	ok,user := models.SelectUserDataByMobile(mp["mobile"].(string))
	if !ok{
		mp["errno"] = models.RECODE_USERERR
		mp["errmsg"] = "账号未注册"
		return
	}
	//比对密码,设置session,返回结果
	if user.Password_hash == mp["password"]{
		mp["errno"] = models.RECODE_OK
		mp["errmsg"] = "密码正确"
		//设置session
		c.SetSession("name",user.Name)
		c.SetSession("id",user.Id)
	}else {
		mp["errno"] = models.RECODE_DBERR
		mp["errmsg"] = "密码错误"
	}

}

func (c *UserController) PostRegisterData() {
	mp := make(map[string]interface{})
	//获取前端的数据
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&mp);err != nil{
		logs.Info(err)
	}

	//插入数据库中,这里的mpreturn用来存储发来的数据
	ok,id	 := models.InsertUserData(mp)
	mpreturn := make(map[string]interface{})
	defer c.sendJSON(mpreturn)

	//后面的mp用来存储发送过去的数据
	if ok == false{
		mpreturn["errno"] = "4002"
		mpreturn["errmsg"] = "注册失败"
	}else{
		//设置session
		c.SetSession("name",mp["mobile"].(string))
		c.SetSession("id",id)
		mpreturn["errno"] = "0"
		mpreturn["errmsg"] = "注册成功"
	}
}

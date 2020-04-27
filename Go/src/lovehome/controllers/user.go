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

	mp["errno"] = models.RECODE_OK
	mp["errmsg"] = models.RecodeText(models.RECODE_OK)
	url := make(map[string]string)
	url["avatar_url"] = path
	mp["data"] = url
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
	ok,user := models.SelectUserData(mp)
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
		c.SetSession("name",user.Password_hash)
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

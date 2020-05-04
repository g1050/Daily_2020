package controllers

import (
	"astaxie/beego"
	"astaxie/beego/logs"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	"lovehome/models"
	"time"
)
type AreaController struct {
	beego.Controller
}

func (c *AreaController) sendJSON(mp map[string]interface{}){
	c.Data["json"]=mp
	c.ServeJSON()
}

func (c *AreaController) GetArea() {

	mp := make(map[string]interface{})
	//打包成json格式发给客户端
	defer c.sendJSON(mp)

	//优化后:利用cache模块和redis,不必每次都从Mysql中取出数据
	//初始化redis数据库
	//key区分不同插入者，类似namespace;conn端口号;dbNum数据库;
	bm, err := cache.NewCache("redis", `{"key":"lovehome","conn":":6379","dbNum":"0"}`)
	if err != nil{
		logs.Error("redis初始化失败",err)
	}
	//尝试从cache中获取数据

	if jsonstr := bm.Get("area");jsonstr!= nil{
		//logs.Debug("Get from cache \n")
		mp["errno"] = models.RECODE_OK
		mp["errmsg"] = models.RecodeText(models.RECODE_OK)
		var area_info interface{}
		json.Unmarshal(jsonstr.([]byte), &area_info)
		mp["data"] = area_info
		return
	}


	//插入数据,第三个参数表示时效性,测试
	/*
	errPut := bm.Put("key1","value1",time.Second*3600)
	if errPut != nil{
		logs.Error("redis插入数据出错",err)
	}
	 */

	//获取数据,测试
	//fmt.Printf("Get data from redis %s\n",bm.Get("key1"))

	//优化前:从数据库中获取数据
	ok,areas := models.SelectAreaDate()

	//将返回的area打包成map
	if ok{//成功获取数据
		mp["errno"] = models.RECODE_OK
		mp["errmsg"] = models.RecodeText(models.RECODE_OK)
		mp["data"] =areas

		//打包成json存入缓存
		json_str,errjson := json.Marshal(areas)

		if errjson != nil{
			logs.Error("打包成json格式失败")
		}else{
			bm.Put("area",json_str,time.Second*3600)
		}
	}else{
		logs.Error("未读取到Area数据")
		mp["errno"] = models.RECODE_DBERR
		mp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	}
}
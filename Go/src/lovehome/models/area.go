package models

import (
	"astaxie/beego/logs"
	"astaxie/beego/orm"
)

func SelectAreaDate()map[string]interface{}{

	//返回一个map数据
	mp := make(map[string]interface{})
	areas := []Area{}
	o := orm.NewOrm()
	num,err := o.QueryTable("area").All(&areas)

	logs.Info("num = ",num)
	if err != nil || num == 0{
		mp["errno"] = 4001
		mp["errmsg"] = "没有查到数据"
		logs.Error(err)
	}else{
		mp["errno"] = 0
		mp["errmsg"] = "OK"
		mp["data"] = areas
	}
	return mp
}


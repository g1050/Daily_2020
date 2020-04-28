package models

import (
	"astaxie/beego/logs"
	"astaxie/beego/orm"
)

func SelectHouseDataByUserID(id int)(bool,[]House){
	houses := []House{}

	o := orm.NewOrm()

	qs := o.QueryTable("house")
	num,err := qs.Filter("user_id",id).All(&houses)

	if err != nil || num == 0{
		logs.Info("house没有数据",err)
		return false,houses
	}

	return true,houses
}
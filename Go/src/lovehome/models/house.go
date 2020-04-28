package models

import (
	"astaxie/beego/logs"
	"astaxie/beego/orm"
)

func InsertHouseData(house House,facility []Facility)(bool,int64){
	o := orm.NewOrm()
	id,err := o.Insert(&house)

	if err != nil{
		return false,id
	}

	//复杂查询、添加
	m2m := o.QueryM2M(&house,"Facilities")
	num,errM2M := m2m.Add(facility)

	if errM2M != nil || num == 0{
		return false,id
	}
	return true,id
}

func SelectHouseDataByUserID(id int)(bool,[]House){
	houses := []House{}

	o := orm.NewOrm()

	qs := o.QueryTable("house")
	num,err := qs.Filter("user__id",id).All(&houses)

	if err != nil || num == 0{
		logs.Info("house没有数据",err)
		return false,houses
	}
	return true,houses
}
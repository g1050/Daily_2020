package models

import (
	"astaxie/beego/orm"
)

func SelectAreaDate()(bool,[]Area){

	//返回一个map数据
	//mp := make(map[string]interface{})
	areas := []Area{}
	o := orm.NewOrm()
	num,err := o.QueryTable("area").All(&areas)
	if err != nil || num == 0{
		return false,areas
	}else {
		return  true,areas
	}

}


package models

import (
	"astaxie/beego/logs"
	"astaxie/beego/orm"
	"fmt"
)

func SelectHouseDataByUserID(id int)(bool,[]House){
	houses := []House{}

	o := orm.NewOrm()

	qs := o.QueryTable("house")
	num,err := qs.Filter("user__id",id).All(&houses)

	if err != nil || num == 0{
		logs.Info("house没有数据",err)
		return false,houses
	}
	fmt.Println("housedata ===== ",houses)
	return true,houses
}
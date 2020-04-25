package models

import (
	"astaxie/beego/logs"
	"astaxie/beego/orm"
)

//M只负责true or false,而不去封装返回的信息
func InsertUserData(mp map[string]interface{})bool{
	o := orm.NewOrm()
	user := User{}
	logs.Info(user)

	user.Password_hash = mp["password"].(string)
	user.Mobile = mp["mobile"].(string)
	user.Name = mp["mobile"].(string)

	_,err := o.Insert(&user)
	if err != nil{
		logs.Error(err)
		return false
	}else{
		return true
	}

}

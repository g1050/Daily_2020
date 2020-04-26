package models

import (
	"astaxie/beego/logs"
	"astaxie/beego/orm"
	"fmt"
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

func SelectUserData(mp map[string]interface{})(bool,User){

	o := orm.NewOrm()
	//将map转化为User结构体
	var user User
	err := o.QueryTable("user").Filter("mobile", mp["mobile"].(string)).One(&user)
	if err == orm.ErrMultiRows {
		return false,user
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		return false,user
		// 没有找到记录
		fmt.Printf("Not row found")
	}
	return true,user

}

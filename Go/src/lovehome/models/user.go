package models

import (
	"astaxie/beego/logs"
	"astaxie/beego/orm"
	"fmt"
)

func UpdateUserData(user User,myfield string)error{
	o := orm.NewOrm()
	_,err := o.Update(&user,myfield)
	return err
}
//M只负责true or false,而不去封装返回的信息
func InsertUserData(mp map[string]interface{})(bool,int64){
	o := orm.NewOrm()
	user := User{}

	user.Password_hash = mp["password"].(string)
	user.Mobile = mp["mobile"].(string)
	//默认name是mobile
	user.Name = mp["mobile"].(string)


	//id是返回的插入后的id
	id,err := o.Insert(&user)

	if err != nil{
		logs.Error(err)
		return false,id
	}else{
		return true,id
	}


}

func SelectUserDataById(id int)(bool,User){

	o := orm.NewOrm()
	//将map转化为User结构体

	//根据mobile查询
	var user User
	err := o.QueryTable("user").Filter("id", id).One(&user)
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

func SelectUserDataByMobile(mobile string)(bool,User){

	o := orm.NewOrm()
	//将map转化为User结构体

	//根据mobile查询
	var user User
	err := o.QueryTable("user").Filter("mobile", mobile).One(&user)
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

package models

import (
	"astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Sex  int8
	User_order []*User_order `orm:"reverse(many)"` // 设置一对多的反向关系
}

type User_order struct {
	Id int
	Descip string
	User *User `orm:"rel(fk)"`//一对多的关系
}

func init() {
	//连接数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8", 30) //最后是一个超时时间
	//注册model
	orm.RegisterModel(new(User),new(User_order))
	//创建表
	orm.RunSyncdb("default",false,true)
}

func PrintUserByORM() {

	o := orm.NewOrm()   //创建ORM对象
	user := User{Id: 3} //创建结构体对象

	o.Read(&user) //将信息读到user中
	fmt.Println(user)
}

//增加数据
func InsertData(){
	o := orm.NewOrm()

	user := User{}
	user.Name = "Chang Yuhang"
	user.Sex = 0

	id,err := o.Insert(&user)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("inser data id = ",id,user)

}

//删除数据
func DeleteData(){
	o := orm.NewOrm()

	user := User{Id:1}

	id,err := o.Delete(&user)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("Delete Data Sucess id = ",id)

}

//修改数据
func UpdateDate(){
	o := orm.NewOrm()

	user := User{Id:3}
	user.Name = "甄姬"
	user.Sex = 4

	id,err := o.Update(&user)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Update data sucessfuly id = ",id)

}

func Read(){
	o := orm.NewOrm()

	user := User{Id:2}

	id := 1
	err := o.Read(&user)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("Read ",id,"data ",user)

}

//Order增加数据
func InsertOrder(){
	o := orm.NewOrm()
	user := User{Id: 2}

	user_order := User_order{User:&user}
	user_order.Descip = "Order1"

	_,err := o.Insert(&user_order)
	if err != nil{
		fmt.Println(err)
		return
	}

}
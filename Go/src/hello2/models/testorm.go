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
}

func init() {
	//连接数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8", 30) //最后是一个超时时间
	//注册model
	orm.RegisterModel(new(User))
	//创建表
	orm.RunSyncdb("default",false,true)
}

func PrintUserByORM() {

	o := orm.NewOrm()   //创建ORM对象
	user := User{Id: 3} //创建结构体对象

	o.Read(&user) //将信息读到user中
	fmt.Println(user)
}

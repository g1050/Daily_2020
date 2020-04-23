package models

import (
	"astaxie/beego/orm"
	"fmt"
)

func init() {
	// orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(MenuModel))
	fmt.Println("Hello World")
}

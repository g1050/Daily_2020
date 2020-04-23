package main

import (
    "fmt"
    "astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
)

type User struct {
    Id int
    Name string 
    sex int8
}

type Student struct{
    Id int
    Name string
    sex int8
}

func init(){

    // 设置默认的数据库
    orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8",30)

    // 注册对象
    orm.RegisterModel(new(Student))

    orm.RunSyncdb("default",false,true)
}   
func main(){

    o := orm.NewOrm()
    student := Student{Name:"Jane"}

    //插入
    id,err := o.Insert(&student)
    fmt.Println(id,err)

    //更新
    student.Name = "Mike"
    num,err := o.Update(&student)
    fmt.Println(num,err)

    //读取一条数据
    read := Student{Id:student.Id}
    err = o.Read(&read)
    fmt.Println(read)

    //删除一条数据
    num,err = o.Delete(&read)
    fmt.Println(num,err)
}


package models

import (
	"astaxie/beego/logs"
	"astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id          int
	Name        string
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"` //设置多对多反向关系
}

func init() {
	//连接Mysql数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/ttms?charset=utf8", 30) //最后是一个超时时间
	//注册model
	orm.RegisterModel(new(User),new(Post),new(Profile),new(Tag))
	//创建表,第二个参数表示如果存在该表是否覆盖
	orm.RunSyncdb("default",false,true)

	logs.Debug("init database .....")
}

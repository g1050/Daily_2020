package main

import (
	"astaxie/beego/logs"
	"astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
	"orm/models"
	_ "orm/models"
	_ "orm/routers"
)



func main() {

	/*一对一关系的定义和查询
	o := orm.NewOrm()
	profile := new(models.Profile)
	profile.Age = 30

	user := new(models.User)
	user.Profile = profile //主表一对一字段直接指向从表
	user.Name = "slene2"

	插入时候先插入从表
	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

	//一对一查询
	user2 := &models.User{Id:1}
	o.Read(user2)
	if user.Profile != nil {
		o.Read(user2.Profile)
	}

	//user2中保存的是指针
	logs.Debug(user2,*user2.Profile)
	*/

	/*
	一对多的定义和查询
	 */

	o := orm.NewOrm()
	user := models.User{Id:11,Name:"王花花"}
	profile := models.Profile{Age:18}
	user.Profile = &profile
	post := models.Post{Title:"总经理"}
	post.User = &user

	/*
	logs.Debug(o.Insert(&profile))
	logs.Debug(o.Insert(&user))
	 */
	logs.Debug(o.Insert(&post))


	var posts []*models.Post
	num, err := o.QueryTable("post").Filter("User", 10).RelatedSel().All(&posts)
	if err == nil {
		fmt.Printf("%d posts read\n", num)
		for _, post := range posts {
			fmt.Printf("Id: %d, UserName: %s, Title: %s\n", post.Id, post.User.Name, post.Title)
		}
	}

	/*
	可以找到所有外键为10的post
	 */
	//默认级联删除,在这里user相当于studio,post相当与seat
	o.Delete(&user)
	beego.Run()
}


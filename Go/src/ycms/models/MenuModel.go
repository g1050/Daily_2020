package models

type MenuModel struct {
	Mid    int `orm:"pk;auto"`
	Parent int
	Name   string `orm:"size(45)"` //varchar(45)
	Seq    int    //排序
	Format string `orm:"size(2048);default({})` //默认是空的json
}

type MenuTree struct {
	MenuModel             //当前节点
	Child     []MenuModel //子节点
}

//确定数据库的表名字
//实现orm中的接口
func (m *MenuModel) TableName() string {
	return "xcms_menu"
}

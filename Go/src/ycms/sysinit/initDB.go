package sysinit

import (
	"astaxie/beego"
	"astaxie/beego/orm"
	_ "ycms/models"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() {

	//连接名称
	dbAlias := beego.AppConfig.String("db_alias")
	//数据库名称
	dbName := beego.AppConfig.String("db_name")
	//用户名
	dbUser := beego.AppConfig.String("db_user")
	//密码
	dbPwd := beego.AppConfig.String("db_pwd")
	//IP
	dbHost := beego.AppConfig.String("db_host")
	//port
	dbPort := beego.AppConfig.String("db_port")
	//charset
	dbCharset := beego.AppConfig.String("db_charset")

	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset, 30)

	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")

	//自动建表
	orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}

}

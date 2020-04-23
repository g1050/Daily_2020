package main

import (
	"astaxie/beego/logs"
	_ "test/routers"

	"github.com/astaxie/beego"
)

func main() {
	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}

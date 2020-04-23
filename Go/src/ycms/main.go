package main

import (
	_ "ycms/routers"
	_ "ycms/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

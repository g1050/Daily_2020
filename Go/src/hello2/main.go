package main

import (
	_ "hello2/routers"
  	_"hello2/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

package main

import (
	"astaxie/beego/context"
	"github.com/astaxie/beego"
	_ "lovehome/routers"
	"net/http"
	"strings"
)

func main() {
	ignore()
	beego.Run()
}


func ignore(){
	//透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	beego.FilterFunc()
	if strings.Index(ctx.Request.URL.Path, "v1/") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+ctx.Request.URL.Path)
}


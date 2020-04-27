package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "lovehome/routers"
	"net/http"
	"strings"
)

func main() {
	ignoreStaticPath()
	beego.Run()
}

type FilterFunc  = func(*context.Context)

/*
func ignore(){
	//透明static
	beego.InsertFilter("/", beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	if strings.Index(ctx.Request.URL.Path, "v1/") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+ctx.Request.URL.Path)
}

*/
func ignoreStaticPath() {

	//透明static
	beego.SetStaticPath("group1/M00/","logs/storage_data/data/")

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)

}

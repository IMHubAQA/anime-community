package router

import (
	"anime-community/controller"

	"github.com/beego/beego/v2/server/web"
)

func Init() {
	initFilter()
	initRouter("/acomm")
	web.Run()
}

func initRouter(prefix string) {
	ns := web.NewNamespace(prefix)
	ns.Router("/test", &controller.TestController{}, "get:Get")
	web.AddNamespace(ns)
}

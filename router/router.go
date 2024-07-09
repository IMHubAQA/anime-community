package router

import (
	"github.com/beego/beego/v2/server/web"

	"anime-community/controller"
)

func Init() {
	initFilter()
	initRouter("/acomm")
	web.Run()
}

func initRouter(prefix string) {
	ns := web.NewNamespace(prefix)
	ns.Router("/test", &controller.TestController{}, "get:Get")

	ns.Router("/post/list", &controller.PostsController{}, "get:List")
	ns.Router("/post/create", &controller.PostsController{}, "post:Create")
	web.AddNamespace(ns)
}

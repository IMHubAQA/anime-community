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

	ns.Router("/post/homepage", &controller.PostsController{}, "get:Homepage")
	ns.Router("/post/create", &controller.PostsController{}, "get:Create")
	web.AddNamespace(ns)
}

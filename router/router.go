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

	ns.Router("/category/list", &controller.CategoryController{}, "get:List")

	ns.Router("/post/list", &controller.PostsController{}, "get:List")
	ns.Router("/post/create", &controller.PostsController{}, "post:Create")
	ns.Router("/post/info", &controller.PostsController{}, "get:Info")
	ns.Router("/post/search", &controller.PostsController{}, "get:Search")

	ns.Router("/comment/list", &controller.CommentController{}, "get:List")
	ns.Router("/comment/create", &controller.CommentController{}, "post:Create")
	web.AddNamespace(ns)
}

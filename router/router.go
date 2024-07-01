package router

import "github.com/beego/beego/v2/server/web"

func Init() {

	web.Run()
}

func initRouter(prefix string) {
	ns := web.NewNamespace(prefix)

	web.AddNamespace(ns)
}

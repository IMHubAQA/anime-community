package main

import (
	_ "anime-community/config"
	"anime-community/router"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("panic :%v", r)
		}
	}()

	logs.Info("%v server is running", web.BConfig.AppName)

	router.Init()
}

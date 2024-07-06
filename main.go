package main

import (
	"context"

	"anime-community/common/helper"
	_ "anime-community/config"
	_ "anime-community/dao/mysql"
	_ "anime-community/dao/redis"
	"anime-community/router"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	defer helper.Recover(context.Background())

	logs.Info(">>[%v] starting......", web.BConfig.AppName)

	run()
}

func run() {
	router.Init()
}

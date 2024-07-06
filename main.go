package main

import (
	"context"

	"anime-community/common/helper"
	"anime-community/common/logs"
	"anime-community/config"
	"anime-community/dao/mysql"
	"anime-community/dao/redis"
	"anime-community/router"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	ctx := logs.NewTraceContext(context.Background())
	defer logs.Sync()
	defer helper.Recover(ctx)

	InitServer(ctx)
}

func InitServer(ctx context.Context) {
	config.Init()
	logs.Init()

	redis.Init(ctx)
	mysql.Init(ctx)

	logs.Infof(ctx, "[%v:%v] starting......", web.BConfig.AppName, web.BConfig.Listen.HTTPPort)

	router.Init()
}

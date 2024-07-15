package main

import (
	"context"
	"flag"

	"github.com/beego/beego/v2/server/web"

	"anime-community/common/helper"
	"anime-community/common/logs"
	"anime-community/config"
	"anime-community/dao/elasticc"
	"anime-community/dao/mysql"
	"anime-community/dao/redis"
	"anime-community/router"
	"anime-community/task"
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
	elasticc.Init(ctx)

	if Tool(ctx) {
		return
	}

	logs.Infof(ctx, "[%v:%v] starting......", web.BConfig.AppName, web.BConfig.Listen.HTTPPort)
	web.BConfig.CopyRequestBody = true
	task.Init()
	router.Init()
}

// 手动执行工具
func Tool(ctx context.Context) bool {
	isTool := flag.Bool("tool", false, "手动执行工具")
	cmd := flag.Int("cmd", 0, "命令id")
	timerId := flag.Int("timerid", 0, "定时任务id")
	flag.Parse()

	if !*isTool {
		return false
	}

	switch *cmd {
	case 1:
		task.ManualExcute(ctx, *timerId)
	default:
		logs.Warnf(ctx, "unkown cmd %v", *cmd)
		flag.Usage()
	}
	return true
}

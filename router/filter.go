package router

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func initFilter() {
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	web.InsertFilter("*", web.BeforeExec, func(ctx *context.Context) {
		logs.Info("method:%v, url:%v, useragent:%v, refer:%v, ip:%v",
			ctx.Request.Method, ctx.Request.URL, ctx.Input.UserAgent(), ctx.Input.Refer(), ctx.Input.IP())
	})
}

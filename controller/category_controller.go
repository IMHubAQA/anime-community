package controller

import (
	"anime-community/common/constants"
	"anime-community/common/helper"
	"anime-community/common/httpc"
	"anime-community/common/logs"
	"anime-community/dao/redis"
)

type CategoryController struct {
	BaseController
}

// 标签列表
func (c *CategoryController) List() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	defer helper.Recover(ctx, func() {
		c.FailJsonResp(constants.ServerError)
	})

	postType, err := c.GetInt("postType")
	if err != nil || postType <= 0 {
		logs.Warnf(ctx, "CategoryController List ParseForm fail. err=%v", err)
		c.FailJsonResp(constants.InvalidParamsError)
		return
	}

	data, err := redis.GetPostsCategoryList(ctx, postType)
	if err != nil {
		c.FailJsonResp(constants.RedisError)
		return
	}

	resp := httpc.NewHttpResult().OkWithData(data).Build()
	c.JsonResp(resp)
}

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
	data, err := redis.GetPostsCategoryList(ctx, c.GetString("postType"))
	if err != nil {
		c.FailJsonResp(constants.RdisError)
	}

	resp := httpc.NewHttpResult().OkWithData(data).Build()
	c.JsonResp(resp)
}

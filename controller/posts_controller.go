package controller

import (
	"anime-community/common/httpc"
	"anime-community/common/logs"
	modelv "anime-community/model/vo"
	"anime-community/service"
)

type PostsController struct {
	BaseController
}

func (c *PostsController) Homepage() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	req := &modelv.PostReq{}
	err := c.ParseForm(req)
	if err != nil {
		logs.Warnf(ctx, "PostsController Homepage ParseForm fail. err=%v", err)
	}

	logs.Infof(ctx, "PostsController Homepage req=%+v", req)

	data, err1 := service.GetHomePage(ctx, req)
	if err1 != nil {
		c.FailJsonResp(err1)
		return
	}

	resp := httpc.NewHttpResult().OkWithData(data).Build()
	c.JsonResp(resp)
}

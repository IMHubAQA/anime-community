package controller

import (
	"anime-community/common/constants"
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
	req := &modelv.PostHomePageReq{}
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

func (c *PostsController) Create() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	req := &modelv.PostCreateReq{}
	req.Uid = c.Ctx.Request.Header.Get("uid")
	req.UToken = c.Ctx.Request.Header.Get("uToken")
	req.PostType = c.Ctx.Request.Header.Get("postType")
	body := c.Ctx.Input.RequestBody

	if req.Uid == "" || req.UToken == "" {
		c.FailJsonResp(constants.InvalidParamsError)
		return
	}

	// 	err := c.ParseForm(req)
	// 	if err != nil {
	// 		logs.Warnf(ctx, "PostsController Homepage ParseForm fail. err=%v", err)
	// 	}

	// 	logs.Infof(ctx, "PostsController Homepage req=%+v", req)

	err := service.CreatePost(ctx, req, body)
	if err != nil {
		c.FailJsonResp(err)
		return
	}
	c.JsonResp(httpc.OkNoData)
}

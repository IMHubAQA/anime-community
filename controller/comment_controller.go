package controller

import (
	"anime-community/common/constants"
	"anime-community/common/helper"
	"anime-community/common/httpc"
	"anime-community/common/logs"
	"anime-community/dao/redis"
	modelv "anime-community/model/vo"
	"anime-community/service"
)

type CommentController struct {
	BaseController
}

// 评论列表
func (c *CommentController) List() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	defer helper.Recover(ctx, func() {
		c.FailJsonResp(constants.ServerError)
	})

	req := &modelv.CommentListReq{}
	err := c.ParseForm(req)
	if err != nil {
		logs.Warnf(ctx, "CommentController List ParseForm fail. err=%v", err)
		c.FailJsonResp(constants.InvalidParamsError)
		return
	}

	req.Init()

	logs.Infof(ctx, "CommentController List req=%+v", req)

	data, err1 := service.GetCommentList(ctx, req)
	if err1 != nil {
		c.FailJsonResp(err1)
		return
	}

	resp := httpc.NewHttpResult().OkWithData(data).Build()
	c.JsonResp(resp)
}

// 发评论
func (c *CommentController) Create() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	defer helper.Recover(ctx, func() {
		c.FailJsonResp(constants.ServerError)
	})
	header, err := modelv.GetAndCheckBaseHeader(ctx, c.Ctx)
	if err != nil {
		c.FailJsonResp(err)
		return
	}

	logs.Infof(ctx, "CommentController Create header=%+v", header)

	// 防止重复提交
	routerLock := redis.CommentCreateRouterLockRedisKey
	if !redis.OnLock(ctx, routerLock, header.Uid) {
		c.FailJsonResp(constants.ServerError)
		return
	}

	err = service.CreateComment(ctx, header, c.Ctx.Input.RequestBody)
	if err != nil {
		c.FailJsonResp(err)
		redis.UnLock(ctx, routerLock, header.Uid)
		return
	}

	c.JsonResp(httpc.OkNoData)
	redis.UnLock(ctx, routerLock, header.Uid)
}

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

type PostsController struct {
	BaseController
}

// 帖子列表
func (c *PostsController) List() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	defer helper.Recover(ctx, func() {
		c.FailJsonResp(constants.ServerError)
	})

	req := &modelv.PostListReq{}
	err := c.ParseForm(req)
	if err != nil {
		logs.Warnf(ctx, "PostsController List ParseForm fail. err=%v", err)
		c.FailJsonResp(constants.InvalidParamsError)
		return
	}

	req.Init()

	logs.Infof(ctx, "PostsController List req=%+v", req)

	data, err1 := service.GetPostList(ctx, req)
	if err1 != nil {
		c.FailJsonResp(err1)
		return
	}

	resp := httpc.NewHttpResult().OkWithData(data).Build()
	c.JsonResp(resp)
}

// 创建帖子
func (c *PostsController) Create() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	defer helper.Recover(ctx, func() {
		c.FailJsonResp(constants.ServerError)
	})
	header, err := modelv.GetAndCheckBaseHeader(ctx, c.Ctx)
	if err != nil {
		c.FailJsonResp(err)
		return
	}

	logs.Infof(ctx, "PostsController Create header=%+v", header)

	// 防止重复提交
	routerLock := redis.PostCreateRouterLockRedisKey
	if !redis.OnLock(ctx, routerLock, header.Uid) {
		c.FailJsonResp(constants.ServerError)
		return
	}

	err = service.CreatePost(ctx, header, c.Ctx.Input.RequestBody)
	if err != nil {
		c.FailJsonResp(err)
		redis.UnLock(ctx, routerLock, header.Uid)
		return
	}

	c.JsonResp(httpc.OkNoData)
	redis.UnLock(ctx, routerLock, header.Uid)
}

// 帖子详情
func (c *PostsController) Info() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	defer helper.Recover(ctx, func() {
		c.FailJsonResp(constants.ServerError)
	})

	req := &modelv.PostInfoReq{}
	err := c.ParseForm(req)
	if err != nil || !req.Check() {
		logs.Warnf(ctx, "PostsController Info ParseForm fail. err=%v", err)
		c.FailJsonResp(constants.InvalidParamsError)
		return
	}

	logs.Infof(ctx, "PostsController Info req=%+v", req)

	data, err1 := service.GetPostById(ctx, req)
	if err1 != nil {
		c.FailJsonResp(err1)
		return
	}

	resp := httpc.NewHttpResult().OkWithData(data).Build()
	c.JsonResp(resp)
}

// 帖子搜索
func (c *PostsController) Search() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	defer helper.Recover(ctx, func() {
		c.FailJsonResp(constants.ServerError)
	})

	req := &modelv.PostSearchReq{}
	err := c.BindJSON(&req)
	if err != nil || !req.Check() {
		logs.Warnf(ctx, "PostsController Search BindJSON fail. err=%v", err)
		c.FailJsonResp(constants.InvalidParamsError)
		return
	}

	logs.Infof(ctx, "PostsController Search req=%+v", req)

	// FIXME :暂时先用list接口作为搜索返回吧，方便联调，数据结构一样，相当于搜索无效。
	data, err1 := service.GetPostList(ctx, &modelv.PostListReq{
		Page:     req.Page,
		PostType: uint64(req.PostType),
		PageSize: 10,
	})

	if err1 != nil {
		c.FailJsonResp(err1)
		return
	}

	resp := httpc.NewHttpResult().OkWithData(data).Build()
	c.JsonResp(resp)
}

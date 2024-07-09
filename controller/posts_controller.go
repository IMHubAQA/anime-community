package controller

import (
	"crypto/sha256"
	"strconv"

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

// 首页
func (c *PostsController) Homepage() {
	ctx := logs.NewTraceContext(c.Ctx.Request.Context())
	defer helper.Recover(ctx, func() {
		c.FailJsonResp(constants.ServerError)
	})

	req := &modelv.PostHomePageReq{}
	err := c.ParseForm(req)
	if err != nil {
		logs.Warnf(ctx, "PostsController Homepage ParseForm fail. err=%v", err)
		c.FailJsonResp(constants.InvalidParamsError)
		return
	}

	req.Init()

	logs.Infof(ctx, "PostsController Homepage req=%+v", req)

	data, err1 := service.GetHomePage(ctx, req)
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

	req := &modelv.PostCreateReq{
		UToken:  c.Ctx.Request.Header.Get("uToken"),
		Sign:    c.Ctx.Request.Header.Get("sign"),
		TimeStr: c.Ctx.Request.Header.Get("timeStr"),
	}
	uid := c.Ctx.Request.Header.Get("uid")
	req.Uid, _ = strconv.Atoi(uid)
	req.PostType, _ = strconv.Atoi(c.Ctx.Request.Header.Get("postType"))

	if req.Uid <= 0 || req.PostType <= 0 || req.UToken == "" {
		c.FailJsonResp(constants.InvalidParamsError)
		return
	}

	if !helper.CheckSign(req.Sign, sha256.New(), uid, req.TimeStr) {
		c.FailJsonResp(constants.InvalidSignError)
		return
	}

	logs.Infof(ctx, "PostsController Create req=%+v", req)

	// 防止重复提交
	routerLock := redis.PostCreateRouterLockRedisKey
	if !redis.OnLock(ctx, routerLock, uid) {
		c.FailJsonResp(constants.ServerError)
		return
	}

	err := service.CreatePost(ctx, req, c.Ctx.Input.RequestBody)
	if err != nil {
		c.FailJsonResp(err)
		redis.UnLock(ctx, routerLock, uid)
		return
	}

	c.JsonResp(httpc.OkNoData)
	redis.UnLock(ctx, routerLock, uid)
}

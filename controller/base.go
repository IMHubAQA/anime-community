package controller

import (
	"anime-community/common/constants"
	"anime-community/common/httpc"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (c *BaseController) JsonResp(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *BaseController) FailJsonResp(err *constants.Error) {
	resp := httpc.NewHttpResult().Fail(err).Build()
	c.Data["json"] = resp
	c.ServeJSON()
}

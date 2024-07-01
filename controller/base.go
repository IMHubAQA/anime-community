package controller

import (
	"anime-community/common/constants"

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
	resp := constants.NewHttpResult().Fail(err).Build()
	c.Data["json"] = resp
	c.ServeJSON()
}

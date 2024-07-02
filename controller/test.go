package controller

import "anime-community/common/httpc"

type TestController struct {
	BaseController
}

func (c *TestController) Get() {
	resp := httpc.NewHttpResult().OkWithData(map[string]interface{}{"key": "value"}).Build()
	c.JsonResp(resp)
}

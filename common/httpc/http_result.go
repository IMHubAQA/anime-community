package httpc

import "anime-community/common/constants"

const (
	_HTTP_CODE_KEY = "code"
	_HTTP_MSG_KEY  = "msg"
	_HTTP_DATA_KEY = "data"
)

var OkNoData = NewHttpResult().Ok().Build()

type HttpResult struct {
	values map[string]interface{}
}

func NewHttpResult() *HttpResult {
	return &HttpResult{
		values: make(map[string]interface{}),
	}
}

func (hr *HttpResult) Build() map[string]interface{} {
	return hr.values
}

func (hr *HttpResult) Append(key string, value interface{}) *HttpResult {
	hr.values[key] = value
	return hr
}

func (hr *HttpResult) Ok() *HttpResult {
	return hr.appendErr(constants.Success)
}

func (hr *HttpResult) OkWithData(data interface{}) *HttpResult {
	hr = hr.appendErr(constants.Success)
	hr.values[_HTTP_DATA_KEY] = data
	return hr
}

func (hr *HttpResult) Fail(err *constants.Error) *HttpResult {
	return hr.appendErr(err)
}

func (hr *HttpResult) appendErr(err *constants.Error) *HttpResult {
	hr.values[_HTTP_CODE_KEY] = err.GetCode()
	hr.values[_HTTP_MSG_KEY] = err.GetMsg()
	return hr
}

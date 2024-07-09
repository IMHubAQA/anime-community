package constants

import "fmt"

type Error struct {
	code int
	msg  string
}

var (
	Success = newError(200, "success")

	commonError        = newError(10000, "") // 自定义msg
	ServerError        = newError(10001, "服务繁忙")
	MysqlError         = newError(10002, "数据库失败")
	RdisError          = newError(10003, "缓存失败")
	NoSupportError     = newError(10004, "功能不支持")
	InvalidParamsError = newError(10005, "非法参数")
	InvalidSignError   = newError(10006, "非法签名")
)

func newError(code int, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) GetCode() int {
	if e == nil {
		return 0
	}
	return e.code
}

func (e *Error) GetMsg() string {
	if e == nil {
		return ""
	}
	return e.msg
}

func (e *Error) IsSuccess() bool {
	if e == nil {
		return true
	}
	return e.code == Success.code
}

func (e *Error) String() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("code=%v, msg=%v", e.code, e.msg)
}

func NewErrorWithMsg(msg string) *Error {
	return &Error{
		code: commonError.code,
		msg:  msg,
	}
}

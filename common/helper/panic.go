package helper

import (
	"context"
	"runtime"

	"github.com/beego/beego/v2/core/logs"
)

func Recover(ctx context.Context, fallBacks ...func()) {
	defer func() {
		if r := recover(); r != nil {
			recoverStat(r)
		}
	}()

	if r := recover(); r != nil {
		recoverStat(r)
		for _, fallBack := range fallBacks {
			if fallBack != nil {
				fallBack()
			}
		}
	}
}

func recoverStat(err interface{}) {
	var buf [2048]byte
	n := runtime.Stack(buf[:], false)
	logs.Error("panic. info=%v, err=%v", n, err)
}

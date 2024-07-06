package helper

import (
	"context"
	"runtime"

	"anime-community/common/logs"
)

func Recover(ctx context.Context, fallBacks ...func()) {
	defer func() {
		if r := recover(); r != nil {
			recoverStat(ctx, r)
		}
	}()

	if r := recover(); r != nil {
		recoverStat(ctx, r)
		for _, fallBack := range fallBacks {
			if fallBack != nil {
				fallBack()
			}
		}
	}
}

func recoverStat(ctx context.Context, err interface{}) {
	var buf [2048]byte
	n := runtime.Stack(buf[:], false)
	logs.Errorf(ctx, "panic. info=%v, err=%v", n, err)
}

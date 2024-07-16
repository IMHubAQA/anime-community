package task

import (
	"context"
	"sync"

	"github.com/beego/beego/v2/task"

	"anime-community/common/helper"
	"anime-community/common/logs"
	"anime-community/task/timer"
)

var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		task.AddTask("ExportPostCategory", task.NewTask("ExportPostCategory", "0 0/1 * * * *",
			func(ctx context.Context) error {
				defer helper.Recover(ctx) //task内部使用了协程，建议加上recover
				return timer.ExportPostCategory(logs.NewTraceContext(ctx))
			}))
		// 执行任务
		task.StartTask()
	})

}

// 手动执行定时任务
func ManualExcute(ctx context.Context, timerId int) {
	switch timerId {
	case 1:
		timer.ExportPostCategory(logs.NewTraceContext(ctx))
	}
}

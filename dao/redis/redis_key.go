package redis

import (
	"fmt"
	"time"
)

const (
	_REDIS_PREFIX     = "acomm"     // 业务前缀
	_REDIS_LOCKPREFIX = "acommlock" // 锁前缀
)

type RedisKey struct {
	Key    string
	Expire time.Duration
}

func (rk *RedisKey) AddSuffix(suffixs ...interface{}) {
	if rk == nil {
		return
	}
	for _, suffix := range suffixs {
		rk.Key = fmt.Sprintf("%v:%v", rk.Key, suffix)
	}
}

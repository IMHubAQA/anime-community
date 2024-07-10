package redis

import (
	"fmt"
	"time"
)

const (
	_REDIS_PREFIX     = "acomm:"     // 业务前缀
	_REDIS_LOCKPREFIX = "acommlock:" // 锁前缀
)

const (
	_DEFAULT_EXPIRE = time.Hour * 24 // 一天
)

var (
	PostCategoryRedisKey = newRedisKey(_REDIS_PREFIX+"pcate", _DEFAULT_EXPIRE*30)
	CommentCountRedisKey = newRedisKey(_REDIS_PREFIX+"ccnt", _DEFAULT_EXPIRE*365)
)

var (
	PostCreateRouterLockRedisKey    = newRedisKey(_REDIS_LOCKPREFIX+"pcrl", time.Second*5)
	CommentCreateRouterLockRedisKey = newRedisKey(_REDIS_LOCKPREFIX+"ccrl", time.Second*5)
)

type RedisKey struct {
	key    string
	expire time.Duration
}

func newRedisKey(key string, expire time.Duration) *RedisKey {
	return &RedisKey{
		key:    key,
		expire: expire,
	}
}

func (rk *RedisKey) GetKey(suffixs ...interface{}) string {
	if rk == nil {
		return ""
	}
	key := rk.key
	for _, suffix := range suffixs {
		key = fmt.Sprintf("%v:%v", key, suffix)
	}
	return key
}

func (rk *RedisKey) GetExpire() time.Duration {
	if rk == nil {
		return 0
	}
	return rk.expire
}

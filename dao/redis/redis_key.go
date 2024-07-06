package redis

import (
	"bytes"
	"time"
)

const (
	_REDIS_PREFIX     = "acomm"     // 业务前缀
	_REDIS_LOCKPREFIX = "acommlock" // 锁前缀
)

const (
	_DEFAULT_EXPIRE = time.Hour * 24 // 一天
)

var PostCategoryRedisKey = newRedisKey(_REDIS_PREFIX+"pcate", _DEFAULT_EXPIRE*30)

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

func (rk *RedisKey) GetKey(suffixs ...string) string {
	if rk == nil {
		return ""
	}
	var bf bytes.Buffer
	bf.WriteString(rk.key)
	for _, suffix := range suffixs {
		bf.WriteString(suffix)
	}
	return bf.String()
}

func (rk *RedisKey) GetExpire() time.Duration {
	if rk == nil {
		return 0
	}
	return rk.expire
}

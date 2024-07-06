package redis

import (
	"context"
	"fmt"
	"time"

	goRedis "github.com/go-redis/redis/v8"
)

const (
	_SINGLE_MUTI_SIZE = 20
)

func splitKeys(curr int, length int) int {
	if curr+_SINGLE_MUTI_SIZE >= length {
		return length
	}
	return curr + _SINGLE_MUTI_SIZE
}

func MutiGet(ctx context.Context, keys []string) (map[string]string, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("empty key")
	}
	offset := 0
	length := len(keys)
	m := make(map[string]string)
	for offset < length {
		netxOffset := splitKeys(offset, length)
		pipe := GetCommunityClient().Pipeline()
		for i := offset; i < netxOffset; i++ {
			pipe.Get(ctx, keys[i])
		}
		offset = netxOffset

		cmders, err := pipe.Exec(ctx)
		if err != nil && err != goRedis.Nil {
			continue
		}
		m := make(map[string]string)
		for i, cmder := range cmders {
			result, err := cmder.(*goRedis.StringCmd).Result()
			if err != nil {
				continue
			}
			m[keys[i]] = result
		}
	}
	return m, nil
}

func MutiSet(ctx context.Context, keys, values []string, expireTime time.Duration) bool {
	if len(keys) != len(values) {
		return false
	}
	offset := 0
	length := len(keys)
	for offset < length {
		netxOffset := splitKeys(offset, length)
		pipe := GetCommunityClient().Pipeline()
		for i := offset; i < netxOffset; i++ {
			pipe.Set(ctx, keys[i], values[i], expireTime)
		}
		offset = netxOffset
		_, err := pipe.Exec(ctx)
		if err != nil {
			return false
		}
	}
	return true
}

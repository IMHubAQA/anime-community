package redis

import (
	"context"
	"fmt"
	"time"

	goRedis "github.com/go-redis/redis/v8"

	"anime-community/common/logs"
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
	logs.Infof(ctx, "MutiGet info. keys=%v", keys)
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
		if err != nil {
			logs.Infof(ctx, "MutiGet fail. err=%v", err)
			continue
		}
		for i, cmder := range cmders {
			result, err := cmder.(*goRedis.StringCmd).Result()
			if err != nil {
				logs.Infof(ctx, "MutiGet fail. err=%v", err)
				continue
			}
			m[keys[i]] = result
		}
	}
	return m, nil
}

func MutiSet(ctx context.Context, keys, values []string, expireTime time.Duration) error {
	if len(keys) != len(values) {
		return fmt.Errorf("len(keys) != len(values)")
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
			return err
		}
	}
	return nil
}

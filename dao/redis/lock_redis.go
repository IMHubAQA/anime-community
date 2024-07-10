package redis

import (
	"context"

	"anime-community/common/logs"
)

func OnLock(ctx context.Context, key *RedisKey, suffixs ...interface{}) bool {
	rediskey := key.GetKey(suffixs...)
	ok, err := GetCommunityClient().SetNX(ctx, rediskey, 1, key.GetExpire()).Result()
	if err != nil || !ok {
		logs.Warnf(ctx, "OnLock fail. key=%v err=%v", rediskey, err)
		return false
	}
	logs.Infof(ctx, "OnLock success. key=%v", rediskey)
	return true
}

func UnLock(ctx context.Context, key *RedisKey, suffixs ...interface{}) bool {
	rediskey := key.GetKey(suffixs...)
	_, err := GetCommunityClient().Del(ctx, rediskey).Result()
	if err != nil {
		logs.Warnf(ctx, "UnLock fail. key=%v err=%v", rediskey, err)
		return false
	}
	logs.Infof(ctx, "UnLock success.%v", rediskey)
	return true
}

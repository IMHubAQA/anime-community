package redis

import (
	"context"

	"anime-community/common/logs"
)

func OnLock(ctx context.Context, key *RedisKey, suffixs ...string) bool {
	err := GetCommunityClient().SetNX(ctx, key.GetKey(suffixs...), 1, key.GetExpire()).Err()
	if err != nil {
		logs.Warnf(ctx, "OnLock fail. err=%v", err)
	}
	return err == nil
}

func UnLock(ctx context.Context, key *RedisKey, suffixs ...string) bool {
	err := GetCommunityClient().Del(ctx, key.GetKey(suffixs...)).Err()
	if err != nil {
		logs.Warnf(ctx, "UnLock fail. err=%v", err)
	}
	return err == nil
}

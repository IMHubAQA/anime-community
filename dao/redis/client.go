package redis

import (
	"anime-community/config"
	"context"
	"sync"
	"time"

	"anime-community/common/logs"

	goRedis "github.com/go-redis/redis/v8"
)

var communityClient *goRedis.ClusterClient
var initOnce sync.Once

func Init(ctx context.Context) {
	initOnce.Do(func() {
		communityClient = newRedisClient()
		logs.Infof(ctx, "load redis success. ")
	})

}

func newRedisClient() *goRedis.ClusterClient {
	if config.GetServerConfig().RedisConfig == nil {
		panic("load redis config fail")
	}
	return goRedis.NewClusterClient(&goRedis.ClusterOptions{
		Addrs:        config.GetServerConfig().RedisConfig.Addr,
		Password:     config.GetServerConfig().RedisConfig.PassWord,
		MinIdleConns: 10,
		PoolSize:     15,

		DialTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		PoolTimeout:  time.Second * 4,
	})
}

func GetCommunityClient() *goRedis.ClusterClient {
	return communityClient
}

package redis

import (
	"anime-community/config"

	goRedis "github.com/go-redis/redis/v8"
)

var communityClient *goRedis.ClusterClient

func init() {
	communityClient = newRedisClient()
}

func newRedisClient() *goRedis.ClusterClient {
	if config.GetServerConfig().RedisConfig == nil {
		panic("load redis config fail")
	}
	return goRedis.NewClusterClient(&goRedis.ClusterOptions{
		Addrs:    config.GetServerConfig().RedisConfig.Addr,
		Password: config.GetServerConfig().RedisConfig.PassWord,
	})
}

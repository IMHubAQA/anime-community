package elasticc

import (
	"context"
	"log"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"

	"anime-community/common/logs"
	"anime-community/config"
)

var communityClient *elasticsearch.Client
var initOnce sync.Once

func Init(ctx context.Context) {
	initOnce.Do(func() {
		communityClient = newElasticClient(ctx)
		// 检查Elasticsearch服务是否可达
		if _, err := communityClient.Ping(); err != nil {
			log.Fatalf("elastic is down: %s", err)
		}
		logs.Infof(ctx, "load elastic success. ")
	})

}

func newElasticClient(ctx context.Context) *elasticsearch.Client {
	log.Println(config.GetServerConfig().ElasticConfig.Addr)
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.GetServerConfig().ElasticConfig.Addr,
	})
	if err != nil {
		logs.Errorf(ctx, "newElasticClient fail. err=%v", err)
	}
	return client
}

func GetCommunityClient() *elasticsearch.Client {
	return communityClient
}

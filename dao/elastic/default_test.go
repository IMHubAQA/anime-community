package elastic

import (
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func AddIndex(ctx context.Context) {
	res, err := esapi.IndexRequest{
		Index:      "test-index",
		DocumentID: "1",
		Body:       strings.NewReader(`{"postId":19,"relayType":2,"relayId":2,"Content":"评论！@#￥￥%~~dsada~"}`),
	}.Do(ctx, GetCommunityClient())
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.IsError() {
		return
	}
}

package redis

import (
	"context"
	"strconv"
)

// 评论加一
func IncrCommentCount(ctx context.Context, replyType, replyId int) error {
	_, err := communityClient.Incr(ctx, CommentCountRedisKey.GetKey(replyType, replyId)).Result()
	return err
}

// 获取评论数
func GetCommentCount(ctx context.Context, replyType, replyId int) (int, error) {
	s, err := communityClient.Get(ctx, CommentCountRedisKey.GetKey(replyType, replyId)).Result()
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num, nil
}

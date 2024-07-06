package mysql

import (
	"anime-community/model"
	"context"
	"fmt"
)

// 获取帖子列表，返回列表，是否最后一页
func GetPostList(ctx context.Context, postType, page, pageSize int) ([]*model.AnimePost, bool, error) {
	tx := communityClient
	resp := []*model.AnimePost{}
	tx.Model(&model.AnimePost{}).
		Where("post_type = ? and status = ?", postType, model.ANIMEPOST_STATUS_VALID).
		Offset(page * pageSize).
		Limit(pageSize).
		Order("create_time desc").
		Find(resp)
	if tx.Error != nil {
		return nil, true, tx.Error
	}

	if len(resp) == 0 {
		return nil, true, fmt.Errorf("empty data")
	}

	return resp, len(resp) < pageSize, nil
}

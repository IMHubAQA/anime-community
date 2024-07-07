package mysql

import (
	"context"
	"fmt"

	modele "anime-community/model/entity"
)

// 获取帖子列表，返回列表，是否最后一页
func GetPostList(ctx context.Context, postType, page, pageSize int) ([]*modele.AnimePost, bool, error) {
	tx := communityClient
	resp := []*modele.AnimePost{}
	tx.Model(&modele.AnimePost{}).
		Where("post_type = ? and status = ?", postType, modele.ANIMEPOST_STATUS_VALID).
		Offset(page * pageSize).
		Limit(pageSize).
		Order("create_time desc").
		Find(&resp)
	if tx.Error != nil {
		return nil, true, tx.Error
	}

	if len(resp) == 0 {
		return nil, true, fmt.Errorf("empty data")
	}

	return resp, len(resp) < pageSize, nil
}

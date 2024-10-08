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

func CreatePost(ctx context.Context, entity *modele.AnimePost) error {
	tx := communityClient.Model(&modele.AnimePost{}).Create(entity)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

// 通过帖子id获取帖子信息
func GetPostById(ctx context.Context, postId uint64) (*modele.AnimePost, error) {
	tx := communityClient
	resp := &modele.AnimePost{}
	tx.Model(&modele.AnimePost{}).
		Where("id = ? and status = ?", postId, modele.ANIMEPOST_STATUS_VALID).
		Find(&resp)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return resp, nil
}

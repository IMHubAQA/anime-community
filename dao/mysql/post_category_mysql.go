package mysql

import (
	"anime-community/model"
	"context"
)

// 获取帖子标签列表，
func GetPostCategoryList(ctx context.Context, postType uint8, offset, limit int) ([]*model.AnimePostCategory, error) {
	tx := communityClient
	resp := []*model.AnimePostCategory{}
	tx.Model(&model.AnimePostCategory{}).
		Where("post_type = ? and status = ?", postType, model.ANIMEPOST_STATUS_VALID).
		Offset(offset).
		Limit(limit).
		Find(resp)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(resp) == 0 {
		return nil, nil
	}

	return resp, nil
}

func GetAllPostCategoryList(ctx context.Context, postType uint8) ([]*model.AnimePostCategory, error) {
	limit := 200
	offset := 0
	res := []*model.AnimePostCategory{}
	for {
		pageRes, err := GetPostCategoryList(ctx, postType, offset, limit)
		if err != nil {
			return nil, err
		}
		res = append(res, pageRes...)
		if len(res) < limit {
			break
		}
		offset += limit
	}
	return res, nil
}

package mysql

import (
	"context"

	modele "anime-community/model/entity"
)

func GetCommentByReplyType(ctx context.Context, replyType, replyId, offset, limit int) ([]*modele.AnimeComment, error) {
	tx := communityClient
	resp := []*modele.AnimeComment{}
	tx.Model(&modele.AnimeComment{}).
		Where("reply_type = ? and reply_id = ? and status = 1", replyType, replyId).
		Offset(offset).
		Limit(limit).
		Order("create_time desc").
		Find(&resp)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(resp) == 0 {
		return nil, nil
	}

	return resp, nil
}

func GetCommentCountByReplyType(ctx context.Context, replyType, replyId, offset, limit int) (int64, error) {
	tx := communityClient
	var count int64
	tx.Model(&modele.AnimeComment{}).
		Where("reply_type = ? and reply_id = ? and status = 1", replyType, replyId).
		Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return count, nil
}

func CreateComment(ctx context.Context, entity *modele.AnimeComment) error {
	tx := communityClient.Model(&modele.AnimeComment{}).Create(entity)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

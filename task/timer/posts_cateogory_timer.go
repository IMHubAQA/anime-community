package timer

import (
	"context"

	"anime-community/common/logs"
	"anime-community/dao/mysql"
	"anime-community/dao/redis"
	"anime-community/model"
)

// 导出帖子标签到Redis
func ExportPostCategory(ctx context.Context) error {
	for postType := range model.ANIMEPOST_TYPE_SET {
		postCategorys, err := mysql.GetAllPostCategoryList(ctx, uint8(postType))
		if err != nil {
			logs.Errorf(ctx, "ExportPostCategory GetAllPostCategoryList fail. postType=%v,err=%v", postType, err)
			continue
		}
		cacheValues := []*model.PostCategoryCache{}
		for _, postCategory := range postCategorys {
			cacheValues = append(cacheValues, model.NewPostCategoryCache(postCategory))
		}
		err = redis.SetPostsCategory(ctx, cacheValues)
		if err != nil {
			logs.Errorf(ctx, "ExportPostCategory SetPostsCategory fail. postType=%v,err=%v", postType, err)
		}
	}
	return nil
}

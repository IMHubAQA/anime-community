package timer

import (
	"context"

	"anime-community/common/logs"
	"anime-community/dao/mysql"
	"anime-community/dao/redis"
	modelc "anime-community/model/cache"
	modele "anime-community/model/entity"
)

// 导出帖子标签到Redis
func ExportPostCategory(ctx context.Context) error {
	for postType := range modele.ANIMEPOST_TYPE_SET {
		postCategorys, err := mysql.GetAllPostCategoryList(ctx, uint8(postType))
		if err != nil {
			logs.Errorf(ctx, "ExportPostCategory GetAllPostCategoryList fail. postType=%v,err=%v", postType, err)
			continue
		}
		cacheValues := []*modelc.PostCategoryCache{}
		for _, postCategory := range postCategorys {
			cacheValues = append(cacheValues, modelc.NewPostCategoryCache(postCategory))
		}
		err = redis.SetPostsCategory(ctx, postType, cacheValues)
		if err != nil {
			logs.Errorf(ctx, "ExportPostCategory SetPostsCategory fail. postType=%v,err=%v", postType, err)
		}
	}
	logs.Infof(ctx, "ExportPostCategory success")
	return nil
}

package service

import (
	"context"

	"github.com/bytedance/sonic"

	"anime-community/common/constants"
	"anime-community/common/logs"
	"anime-community/dao/mysql"
	"anime-community/dao/redis"
	modele "anime-community/model/entity"
	modelv "anime-community/model/vo"
)

func GetHomePage(ctx context.Context, req *modelv.PostReq) (*modelv.PostResp, *constants.Error) {
	switch req.PostType {
	case modele.ANIMEPOST_TYPE_NORMAL:
		return getHomePageNormal(ctx, req)
	case modele.ANIMEPOST_TYPE_MAKEUP:
		return nil, constants.NoSupportError
		// return getHomePageMakeup(ctx, req)
	default:
		return nil, constants.NoSupportError
	}
}

// 普通贴
func getHomePageNormal(ctx context.Context, req *modelv.PostReq) (*modelv.PostResp, *constants.Error) {
	posts, isLastPage, err := mysql.GetPostList(ctx, int(req.PostType), int(req.Page), 10)
	if err != nil {
		logs.Errorf(ctx, "GetHomePage GetPostList fail. err=%v", err)
		return nil, constants.MysqlError
	}
	categoryIds := []int{}
	appendInfo := make([]*modele.AnimePostAppendNormal, len(posts))
	for i, post := range posts {
		appendInfo[i] = &modele.AnimePostAppendNormal{}
		err := sonic.Unmarshal([]byte(post.AppendInfo), appendInfo[i])
		if err != nil {
			continue
		}
		categoryIds = append(categoryIds, appendInfo[i].Category...)
	}

	categorym, err := redis.GetPostsCategory(ctx, categoryIds)
	if err != nil {
		logs.Warnf(ctx, "GetHomePage GetPostsCategory fail. err=%v", err)
	}

	datas := []*modelv.PostData{}
	for i, post := range posts {
		data := &modelv.PostData{
			PostId:      uint64(post.Id),
			PostType:    req.PostType,
			PostTitle:   post.Title,
			PostContent: post.Content,
			Author: &modelv.PostDataAuthor{
				Uid:  uint64(post.UserId),
				Name: "哈哈",
			},
			CreateTime: uint64(post.CreateTime),
		}

		for _, categoryId := range appendInfo[i].Category {
			if categoryInfo, ok := categorym[categoryId]; ok {
				data.PostCategory = append(data.PostCategory, &modelv.PostDataCategory{
					Id:   uint64(categoryInfo.Id),
					Name: categoryInfo.Name,
				})
			}
		}
		datas = append(datas, data)
	}

	return &modelv.PostResp{
		IsLastPage: isLastPage,
		PostList:   datas,
	}, nil
}

// func getHomePageMakeup(ctx context.Context, req *modelv.PostReq) (*modelv.PostResp, *constants.Error)

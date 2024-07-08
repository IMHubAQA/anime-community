package service

import (
	"context"

	"github.com/bytedance/sonic"

	"anime-community/common/constants"
	"anime-community/common/logs"
	"anime-community/dao/mysql"
	"anime-community/dao/redis"
	modelc "anime-community/model/cache"
	modele "anime-community/model/entity"
	modelv "anime-community/model/vo"
)

func GetHomePage(ctx context.Context, req *modelv.PostHomePageReq) (*modelv.PostHomePageResp, *constants.Error) {
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
func getHomePageNormal(ctx context.Context, req *modelv.PostHomePageReq) (*modelv.PostHomePageResp, *constants.Error) {
	posts, isLastPage, err := mysql.GetPostList(ctx, int(req.PostType), int(req.Page), 10)
	if err != nil {
		logs.Errorf(ctx, "GetHomePage GetPostList fail. err=%v", err)
		return nil, constants.MysqlError
	}

	appendInfom, categorym := getPostInfo(ctx, posts)
	datas := []*modelv.PostData{}

	for _, post := range posts {
		data := &modelv.PostData{
			PostId:      uint64(post.Id),
			PostType:    req.PostType,
			PostTitle:   post.Title,
			PostContent: post.Content,
			CreateTime:  uint64(post.CreateTime),
		}
		buildPostAuthor(ctx, data, post.UserId)
		buildPostAppendInfo(ctx, data, appendInfom, categorym)
		datas = append(datas, data)
	}

	return &modelv.PostHomePageResp{
		IsLastPage: isLastPage,
		PostList:   datas,
	}, nil
}

// 获取帖子依赖的一些信息
func getPostInfo(
	ctx context.Context,
	posts []*modele.AnimePost,
) (
	map[int64]*modele.AnimePostAppendNormal,
	map[int]*modelc.PostCategoryCache,
) {
	categoryIds := []int{}
	appendInfom := make(map[int64]*modele.AnimePostAppendNormal)
	for _, post := range posts {
		appendInfo := &modele.AnimePostAppendNormal{}
		err := sonic.Unmarshal([]byte(post.AppendInfo), appendInfo)
		if err != nil {
			continue
		}
		appendInfom[post.Id] = appendInfo
		categoryIds = append(categoryIds, appendInfo.Category...)
	}

	categorym, err := redis.GetPostsCategory(ctx, categoryIds)
	if err != nil {
		logs.Warnf(ctx, "GetHomePage GetPostsCategory fail. err=%v", err)
	}
	//TODO:用户信息
	return appendInfom, categorym
}

// 生成用户信息
func buildPostAuthor(
	ctx context.Context,
	data *modelv.PostData,
	userId int64,
) {
	data.Author = &modelv.PostDataAuthor{
		Uid:  uint64(userId),
		Name: "哈哈",
	}
}

// 生成附加信息，目前主要是标签
func buildPostAppendInfo(
	ctx context.Context,
	data *modelv.PostData,
	appendInfom map[int64]*modele.AnimePostAppendNormal,
	categorym map[int]*modelc.PostCategoryCache,
) {
	if appendInfo, ok := appendInfom[int64(data.PostId)]; ok {
		for _, categoryId := range appendInfo.Category {
			if categoryInfo, ok := categorym[categoryId]; ok {
				data.PostCategory = append(data.PostCategory, &modelv.PostDataCategory{
					Id:   uint64(categoryInfo.Id),
					Name: categoryInfo.Name,
				})
			}
		}
	}
}

// func getHomePageMakeup(ctx context.Context, req *modelv.PostReq) (*modelv.PostResp, *constants.Error)

func CreatePost(ctx context.Context, req *modelv.PostCreateReq, body []byte) *constants.Error {
	return nil
}

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
		return getHomePageMakeup(ctx, req)
	default:
		return nil, constants.NoSupportError
	}
}

// 普通贴
func getHomePageNormal(ctx context.Context, req *modelv.PostHomePageReq) (*modelv.PostHomePageResp, *constants.Error) {
	posts, isLastPage, err := mysql.GetPostList(ctx, int(req.PostType), int(req.Page-1), 10)
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
		data = buildPostAuthor(ctx, data, post.UserId)
		data = buildPostAppendInfo(ctx, data, appendInfom, categorym)
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
) *modelv.PostData {
	data.Author = &modelv.PostDataAuthor{
		Uid:  uint64(userId),
		Name: "哈哈",
	}
	return data
}

// 生成附加信息，目前主要是标签
func buildPostAppendInfo(
	ctx context.Context,
	data *modelv.PostData,
	appendInfom map[int64]*modele.AnimePostAppendNormal,
	categorym map[int]*modelc.PostCategoryCache,
) *modelv.PostData {
	appendInfo, ok := appendInfom[int64(data.PostId)]
	if !ok {
		return data
	}
	logs.Infof(ctx, "%v  %v", appendInfo, categorym)
	dAppendInfo := &modelv.PostAppendInfo{}
	for _, categoryId := range appendInfo.Category {
		if categoryInfo, ok := categorym[categoryId]; ok {
			dAppendInfo.Category = append(dAppendInfo.Category, &modelv.PostDataCategory{
				Id:   uint64(categoryInfo.Id),
				Name: categoryInfo.Name,
			})
		}
	}
	data.AppendInfo = dAppendInfo

	return data
}

// 获取帖子依赖的一些信息
func getPostInfoMakeUp(
	ctx context.Context,
	posts []*modele.AnimePost,
) (
	map[int64]*modele.AnimePostAppendMakeup,
	map[int]*modelc.PostCategoryCache,
) {
	categoryIds := []int{}
	appendInfom := make(map[int64]*modele.AnimePostAppendMakeup)
	for _, post := range posts {
		appendInfo := &modele.AnimePostAppendMakeup{}
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

// 生成附加信息，
func buildPostAppendInfoMakeup(
	ctx context.Context,
	data *modelv.PostData,
	appendInfom map[int64]*modele.AnimePostAppendMakeup,
	categorym map[int]*modelc.PostCategoryCache,
) *modelv.PostData {
	appendInfo, ok := appendInfom[int64(data.PostId)]
	if !ok {
		return data
	}
	dAppendInfo := &modelv.PostAppendInfo{
		OnDoor: appendInfo.OnDoor,
		Price:  appendInfo.Price,
		Locate: appendInfo.Locate,
	}
	for _, categoryId := range appendInfo.Category {
		if categoryInfo, ok := categorym[categoryId]; ok {
			dAppendInfo.Category = append(dAppendInfo.Category, &modelv.PostDataCategory{
				Id:   uint64(categoryInfo.Id),
				Name: categoryInfo.Name,
			})
		}
	}
	data.AppendInfo = dAppendInfo
	return data
}

func getHomePageMakeup(ctx context.Context, req *modelv.PostHomePageReq) (*modelv.PostHomePageResp, *constants.Error) {
	posts, isLastPage, err := mysql.GetPostList(ctx, int(req.PostType), int(req.Page-1), 10)
	if err != nil {
		logs.Errorf(ctx, "getHomePageMakeup GetPostList fail. err=%v", err)
		return nil, constants.MysqlError
	}

	appendInfom, categorym := getPostInfoMakeUp(ctx, posts)
	datas := []*modelv.PostData{}

	for _, post := range posts {
		data := &modelv.PostData{
			PostId:      uint64(post.Id),
			PostType:    req.PostType,
			PostTitle:   post.Title,
			PostContent: post.Content,
			CreateTime:  uint64(post.CreateTime),
		}
		data = buildPostAuthor(ctx, data, post.UserId)
		data = buildPostAppendInfoMakeup(ctx, data, appendInfom, categorym)
		datas = append(datas, data)
	}

	return &modelv.PostHomePageResp{
		IsLastPage: isLastPage,
		PostList:   datas,
	}, nil
}

func CreatePost(ctx context.Context, req *modelv.PostCreateReq, body []byte) *constants.Error {
	//TODO: AUTH
	bodyData := &modelv.PostCreateBody{}
	err := sonic.Unmarshal(body, bodyData)
	if err != nil {
		logs.Errorf(ctx, "CreatePost Unmarshal fail. body=%v err=%v", string(body), err)
		return constants.InvalidParamsError
	}
	if !bodyData.Check(req.PostType) {
		return constants.InvalidParamsError
	}

	entity := &modele.AnimePost{
		PostType: uint8(req.PostType),
		UserId:   int64(req.Uid),
		Title:    bodyData.Title,
		Content:  bodyData.Content,
		Status:   modele.ANIMEPOST_STATUS_VALID,
	}
	// media
	if len(bodyData.Media) > 0 {
		media, err := sonic.Marshal(bodyData.Media)
		if err != nil {
			return constants.InvalidParamsError
		}
		entity.Media = string(media)
	}
	// appendInfo
	if ainfo := getAppendInfo(ctx, req.PostType, bodyData); ainfo != nil {
		appendInfo, err := sonic.Marshal(ainfo)
		if err != nil {
			return constants.InvalidParamsError
		}
		entity.AppendInfo = string(appendInfo)
	}

	err = mysql.CreatePost(ctx, entity)
	if err != nil {
		logs.Errorf(ctx, "CreatePost insertDb fail. err=%v", err)
		return constants.MysqlError
	}

	return nil
}

func getAppendInfo(ctx context.Context, postType int, body *modelv.PostCreateBody) interface{} {
	if body == nil {
		return nil
	}
	switch postType {
	case modele.ANIMEPOST_TYPE_NORMAL:
		return modele.AnimePostAppendNormal{
			Category: body.Category,
		}
	case modele.ANIMEPOST_TYPE_MAKEUP:
		return modele.AnimePostAppendMakeup{
			Category: body.Category,
			OnDoor:   body.OnDoor,
			Price:    body.Price,
			Locate:   body.Locate,
		}
	default:
		return nil
	}
}

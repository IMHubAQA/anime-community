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

// 获取贴子列表
func GetPostList(ctx context.Context, req *modelv.PostListPistageReq) (*modelv.PostListPageResp, *constants.Error) {
	posts, isLastPage, err := mysql.GetPostList(ctx, int(req.PostType), int(req.Page-1), 10)
	if err != nil {
		logs.Errorf(ctx, "GetHomePage GetPostList fail. err=%v", err)
		return nil, constants.MysqlError
	}

	categorym := getCategoryInfo(ctx, posts)

	datas := []*modelv.PostData{}
	for _, post := range posts {
		data := &modelv.PostData{
			PostId:      uint64(post.Id),
			PostType:    req.PostType,
			PostTitle:   post.Title,
			PostContent: post.Content,
			OnDoor:      int(post.Ondoor),
			Price:       float64(post.Price),
			Location:    post.Location,
			CreateTime:  uint64(post.CreateTime),
		}
		data = buildPostAuthor(ctx, data, post.UserId)
		data.Category = categorym[int64(data.PostId)]
		datas = append(datas, data)
	}

	return &modelv.PostListPageResp{
		IsLastPage: isLastPage,
		PostList:   datas,
	}, nil
}

// 获取标签信息
func getCategoryInfo(
	ctx context.Context,
	posts []*modele.AnimePost,
) map[int64][]*modelv.PostDataCategory {
	categoryIds := []int{}
	postcIdMap := make(map[int64][]int)
	for _, post := range posts {
		category := []int{}
		err := sonic.Unmarshal([]byte(post.Category), &category)
		if err != nil {
			continue
		}
		postcIdMap[post.Id] = category
		categoryIds = append(categoryIds, category...)
	}

	categorym, err := redis.GetPostsCategory(ctx, categoryIds)
	if err != nil {
		logs.Warnf(ctx, "GetHomePage GetPostsCategory fail. err=%v", err)
		return nil
	}
	postCategoryInfo := make(map[int64][]*modelv.PostDataCategory)
	for postId, ids := range postcIdMap {
		postCategoryInfo[postId] = []*modelv.PostDataCategory{}
		for _, id := range ids {
			if info, ok := categorym[id]; ok {
				postCategoryInfo[postId] = append(postCategoryInfo[postId], &modelv.PostDataCategory{
					Id:   uint64(info.Id),
					Name: info.Name,
				})
			}
		}
	}
	return postCategoryInfo
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

func CreatePost(ctx context.Context, req *modelv.PostCreateReq, body []byte) *constants.Error {
	//TODO: AUTH
	bodyData := &modelv.PostCreateBody{}
	err := sonic.Unmarshal(body, bodyData)
	if err != nil {
		logs.Errorf(ctx, "CreatePost Unmarshal fail. body=%v err=%v", string(body), err)
		return constants.InvalidParamsError
	}
	if !bodyData.Check() {
		return constants.InvalidParamsError
	}

	entity := &modele.AnimePost{
		PostType: uint8(bodyData.PostType),
		UserId:   int64(req.Uid),
		Title:    bodyData.Title,
		Content:  bodyData.Content,
		Status:   modele.ANIMEPOST_STATUS_VALID,
		Ondoor:   uint8(bodyData.OnDoor),
		Price:    int64(bodyData.Price),
		Location: bodyData.Location,
	}
	// category
	if len(bodyData.Category) > 0 {
		category, err := sonic.Marshal(bodyData.Category)
		if err != nil {
			return constants.InvalidParamsError
		}
		entity.Category = string(category)
	}
	// media
	if len(bodyData.Media) > 0 {
		media, err := sonic.Marshal(bodyData.Media)
		if err != nil {
			return constants.InvalidParamsError
		}
		entity.Media = string(media)
	}
	err = mysql.CreatePost(ctx, entity)
	if err != nil {
		logs.Errorf(ctx, "CreatePost insertDb fail. err=%v", err)
		return constants.MysqlError
	}

	return nil
}

package service

import (
	"context"
	"encoding/json"
	"time"

	"anime-community/common/constants"
	"anime-community/common/logs"
	commservice "anime-community/common/service"
	"anime-community/dao/mysql"
	"anime-community/dao/redis"
	modele "anime-community/model/entity"
	modelv "anime-community/model/vo"
)

func buildPostData(ctx context.Context, post *modele.AnimePost) *modelv.PostData {
	data := &modelv.PostData{
		PostId:      uint64(post.Id),
		PostType:    uint64(post.PostType),
		PostTitle:   post.Title,
		PostContent: post.Content,
		OnDoor:      int(post.Ondoor),
		Price:       float64(post.Price),
		Location:    post.Location,
		CreateTime:  uint64(post.CreateTime),
	}
	media := []*modelv.MediaData{}
	if err := json.Unmarshal([]byte(post.Media), &media); err == nil {
		data.Media = media
	}
	if commentCnt, err := redis.GetCommentCount(ctx, modele.ANIMECOMMENT_REPLYTYPE_POST, int(data.PostId)); err == nil {
		data.ReplyCnt = uint64(commentCnt)
	}
	return data
}

// 通过帖子id获取帖子信息
func GetPostById(ctx context.Context, req *modelv.PostInfoReq) (*modelv.PostData, *constants.Error) {
	post, err := mysql.GetPostById(ctx, req.PostId)
	if err != nil {
		logs.Errorf(ctx, "GetPostById GetPostById fail. err=%v", err)
		return nil, constants.MysqlError
	}

	categorym, userInfom := getPostextraData(ctx, []*modele.AnimePost{post})

	data := buildPostData(ctx, post)
	data.Category = categorym[int64(post.Id)]
	data.Author = buildAuthor(ctx, userInfom[int(post.UserId)])
	return data, nil
}

// 获取贴子列表
func GetPostList(ctx context.Context, req *modelv.PostListReq) (*modelv.PostListResp, *constants.Error) {
	posts, isLastPage, err := mysql.GetPostList(ctx, int(req.PostType), int(req.Page-1), req.PageSize)
	if err != nil {
		logs.Errorf(ctx, "GetPostList GetPostList fail. err=%v", err)
		return nil, constants.MysqlError
	}

	categorym, userInfom := getPostextraData(ctx, posts)

	datas := []*modelv.PostData{}
	for _, post := range posts {
		data := buildPostData(ctx, post)
		data.Category = categorym[int64(post.Id)]
		data.Author = buildAuthor(ctx, userInfom[int(post.UserId)])
		datas = append(datas, data)
	}

	return &modelv.PostListResp{
		IsLastPage: isLastPage,
		PostList:   datas,
	}, nil
}

// 获取标签信息
func getPostextraData(
	ctx context.Context,
	posts []*modele.AnimePost,
) (map[int64][]*modelv.PostDataCategory, map[int]*commservice.UserData) {
	categoryIds := []int{}
	postcIdMap := make(map[int64][]int)
	// FIXME: make data split
	userIds := []int64{}
	for _, post := range posts {
		category := []int{}
		err := json.Unmarshal([]byte(post.Category), &category)
		if err != nil {
			continue
		}
		postcIdMap[post.Id] = category
		categoryIds = append(categoryIds, category...)
		userIds = append(userIds, post.UserId)
	}

	categorym, err := redis.GetPostsCategory(ctx, categoryIds)
	if err != nil {
		logs.Warnf(ctx, "GetHomePage GetPostsCategory fail. err=%v", err)
		return nil, nil
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
	userInfom, err := commservice.MGetUserInfom(ctx, &commservice.MGetUserReq{UserIds: userIds}, time.Second)
	if err != nil {
		logs.Warnf(ctx, "GetHomePage GetPostsCategory fail. err=%v", err)
	}
	return postCategoryInfo, userInfom
}

// 生成用户信息
func buildAuthor(
	ctx context.Context,
	userInfo *commservice.UserData,
) *modelv.AuthorData {
	if userInfo != nil {
		return nil
	}
	return &modelv.AuthorData{
		Uid:  uint64(userInfo.ID),
		Name: userInfo.NickName,
		Icon: userInfo.AvatarURL,
	}
}

func CreatePost(ctx context.Context, req *modelv.BaseHeader, body []byte) *constants.Error {
	bodyData := &modelv.PostCreateBody{}
	err := json.Unmarshal(body, bodyData)
	if err != nil {
		logs.Errorf(ctx, "CreatePost Unmarshal fail. body=%v err=%v", string(body), err)
		return constants.NewErrorWithMsg(err.Error())
	}
	if err := bodyData.Check(); err != nil {
		return constants.NewErrorWithMsg(err.Error())
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
		category, err := json.Marshal(bodyData.Category)
		if err != nil {
			return constants.NewErrorWithMsg(err.Error())
		}
		entity.Category = string(category)
	}
	// media
	if len(bodyData.Media) > 0 {
		media, err := json.Marshal(bodyData.Media)
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

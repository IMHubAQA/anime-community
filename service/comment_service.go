package service

import (
	"context"
	"encoding/json"

	"anime-community/common/constants"
	"anime-community/common/helper"
	"anime-community/common/logs"
	"anime-community/dao/mysql"
	"anime-community/dao/redis"
	modele "anime-community/model/entity"
	modelv "anime-community/model/vo"
)

// 评论列表
func GetCommentList(ctx context.Context, req *modelv.CommentListReq) (*modelv.CommentListResp, *constants.Error) {
	commentList, err := mysql.GetCommentByReplyType(ctx, int(req.ReplyType), int(req.ReplyId), int(req.Page)-1, req.PageSize)
	if err != nil {
		logs.Errorf(ctx, "GetCommentList GetCommentByReplyType. err=%v", err)
		return nil, constants.MysqlError
	}
	commentRespList := []*modelv.CommentData{}
	for _, comment := range commentList {
		data := &modelv.CommentData{
			CommentId: uint64(comment.Id),
			Content:   comment.Content,
			Author: &modelv.AuthorData{
				Uid:  1,
				Name: "哈哈哈哈",
			},
			CreateTime: uint64(comment.CreateTime),
			PostId:     uint64(comment.PostId),
		}
		if req.ReplyType == modele.ANIMECOMMENT_REPLYTYPE_POST {
			if commentCnt, err := redis.GetCommentCount(ctx, modele.ANIMECOMMENT_REPLYTYPE_COMMENT, int(comment.Id)); err == nil {
				data.ReplyCnt = uint64(commentCnt)
			}
		}
		commentRespList = append(commentRespList, data)
	}
	return &modelv.CommentListResp{
		CommentList: commentRespList,
		IsLastPage:  len(commentList) < req.PageSize,
	}, nil
}

// 创建评论
func CreateComment(ctx context.Context, req *modelv.BaseHeader, body []byte) *constants.Error {
	bodyData := &modelv.CommentCreateJsonBody{}
	err := json.Unmarshal(body, bodyData)
	if err != nil {
		logs.Errorf(ctx, "CreateComment Unmarshal fail. body=%v err=%v", string(body), err)
		return constants.NewErrorWithMsg("invalid jsondata")
	}
	if !bodyData.Check() {
		return constants.NewErrorWithMsg("invalid jsondata")
	}
	entity := &modele.AnimeComment{
		Content:      bodyData.Content,
		UserId:       int64(req.Uid),
		PostId:       int64(bodyData.PostId),
		ReplyType:    uint8(bodyData.RelayType),
		ReplyId:      int64(bodyData.RelayId),
		TargetUserId: int64(bodyData.TargetUid),
		Status:       modele.ANIMECOMMENT_STATUS_VALID,
	}
	err = mysql.CreateComment(ctx, entity)
	if err != nil {
		logs.Errorf(ctx, "CreateComment insertDb fail. err=%v", err)
		return constants.MysqlError
	}
	// 增加评论数
	go func() {
		defer helper.Recover(ctx)
		if err := redis.IncrCommentCount(ctx, int(bodyData.RelayType), int(bodyData.RelayId)); err != nil {
			logs.Errorf(ctx, "CreateComment IncrCommentCount fail. err=%v", err)
		}
	}()
	return nil
}

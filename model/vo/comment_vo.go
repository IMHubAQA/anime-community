package modelv

import modele "anime-community/model/entity"

type CommentListReq struct {
	Page uint64 `form:"page"`
	// PostId    uint64 `form:"postId"`    // 帖子id
	RelayType uint64 `form:"relayType"` // 回复类型：1: 帖子，2：评论
	RelayId   uint64 `form:"relayId"`   // 回复类型id

	PageSize int `form:"-"`
}

func (r *CommentListReq) Init() {
	if r == nil {
		r = &CommentListReq{}
	}
	if r.Page < 1 {
		r.Page = 1
	}
	r.PageSize = 10
}

type CommentListResp struct {
	IsLastPage  bool           `json:"isLastPage"`  // 是否最后一页
	CommentList []*CommentData `json:"CommentList"` // 帖子列表
}

type CommentData struct {
	PostId       uint64      `json:"postId"`                 // 帖子id
	CommentId    uint64      `json:"commentId"`              // 评论id
	Content      string      `json:"content"`                // 回复内容
	Author       *AuthorData `json:"author"`                 // 回复者
	TargetAuthor *AuthorData `json:"targetAuthor,omitempty"` // 回复目标用户
	CreateTime   uint64      `json:"createTime"`             // 回复时间
	ReplyCnt     uint64      `json:"replyCnt"`               // 回复数
}

type CommentCreateJsonBody struct {
	PostId    uint64 `json:"postId"`    // 帖子id
	RelayType uint64 `json:"relayType"` // 1: 帖子，2：评论
	RelayId   uint64 `json:"relayId"`
	Content   string `json:"content"`
	TargetUid int    `json:"targetUid"`
}

func (b *CommentCreateJsonBody) Check() bool {
	if b == nil {
		return false
	}
	if (b.RelayType != modele.ANIMECOMMENT_REPLYTYPE_POST && b.RelayType != modele.ANIMECOMMENT_REPLYTYPE_COMMENT) ||
		b.Content == "" {
		return false
	}
	return true
}

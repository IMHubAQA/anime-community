package modelv

import modele "anime-community/model/entity"

type PostListReq struct {
	// Uid      uint64 `form:"uid"`
	Page     uint64 `form:"page"`
	PostType uint64 `form:"postType"`
	// Category uint64 `form:"category"` // 分类筛选
	PageSize int `form:"-"`
}

func (req *PostListReq) Init() {
	if req == nil {
		req = &PostListReq{}
	}
	if req.Page == 0 {
		req.Page = 1
	}
	req.PageSize = 10
}

// func (req *PostReq) Check()

type PostListResp struct {
	IsLastPage bool        `json:"isLastPage"` // 是否最后一页
	PostList   []*PostData `json:"postList"`   // 帖子列表
}

type PostData struct {
	PostId      uint64              `json:"postId"`             // 帖子id
	PostType    uint64              `json:"postType"`           // 帖子类型
	PostTitle   string              `json:"postTitle"`          // 标题
	PostContent string              `json:"postContent"`        // 内容
	Media       []*MediaData        `json:"media,omitempty"`    // 图片、视频
	Author      *AuthorData         `json:"author"`             // 作者
	LikeCnt     uint64              `json:"LikeCnt"`            // 点赞
	ReplyCnt    uint64              `json:"ReplyCnt"`           // 回复
	CollectCnt  uint64              `json:"collectCnt"`         // 收藏
	CreateTime  uint64              `json:"createTime"`         // 时间
	Category    []*PostDataCategory `json:"category"`           // 标签
	OnDoor      int                 `json:"onDoror,omitempty"`  // 是否可上门
	Price       float64             `json:"price,omitempty"`    // 价格
	Location    string              `json:"location,omitempty"` // 地址
}

type PostDataCategory struct {
	Id   uint64 `json:"id"`   // 标签id
	Name string `json:"name"` // 标签名称
}

type PostCreateBody struct {
	PostType int          `json:"postType"` // 帖子类型
	Title    string       `json:"title"`    // 标题
	Content  string       `json:"content"`  // 内容
	Media    []*MediaData `json:"media"`    // 图片/视频
	Category []int        `json:"category"` // 标签id列表
	OnDoor   int          `json:"onDoror"`  // 0 : 不可上门，1：可以
	Price    float64      `json:"price"`    // 价格
	Location string       `json:"location"` // 地址
}

func (b *PostCreateBody) Check() bool {
	if b == nil {
		return false
	}
	if _, ok := modele.ANIMEPOST_TYPE_SET[b.PostType]; !ok {
		return false
	}
	if b.Title == "" || b.Content == "" || len(b.Category) == 0 {
		return false
	}
	return true
}

type PostInfoReq struct {
	PostId uint64 `form:"postId"`
}

func (req *PostInfoReq) Check() bool {
	if req == nil {
		return false
	}
	return req.PostId > 0
}

type PostSearchReq struct {
	Keyword  string `json:"keyword"` // 关键词
	Page     uint64 `json:"page"`
	PostType int    `json:"postType"` // 帖子类型
}

func (b *PostSearchReq) Check() bool {
	if b == nil {
		return false
	}
	if b.Keyword == "" {
		return false
	}
	return true
}

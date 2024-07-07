package modelv

type PostReq struct {
	Uid      uint64 `form:"uid"`
	Page     uint64 `form:"page"`
	PostType uint64 `form:"postType"`
	Category uint64 `form:"category"`
}

// func (req *PostReq) Init()
// func (req *PostReq) Check()

type PostResp struct {
	IsLastPage bool        `json:"isLastPage"` // 是否最后一页
	PostList   []*PostData `json:"postList"`   // 帖子列表
}

type PostData struct {
	PostId       uint64              `json:"postId"`          // 帖子id
	PostType     uint64              `json:"postType"`        // 帖子类型
	PostTitle    string              `json:"postTitle"`       // 标题
	PostContent  string              `json:"postContent"`     // 内容
	Media        []*PostDataMedia    `json:"media,omitempty"` // 图片、视频
	Author       *PostDataAuthor     `json:"author"`          // 作者
	PostCategory []*PostDataCategory `json:"postCategory"`    // 标签
	LikeCnt      uint64              `json:"LikeCnt"`         // 点赞
	ReplyCnt     uint64              `json:"ReplyCnt"`        // 回复
	CollectCnt   uint64              `json:"collectCnt"`      // 收藏
	CreateTime   uint64              `json:"createTime"`      // 时间
}

type PostDataMedia struct {
	MType    int    `json:"mType"`  // 1:图片；2:视频
	PicUrl   string `json:"picUrl"` // 如果是视频，该字段为视频封面地址
	VideoUrl string `json:"videoUrl"`
}

type PostDataAuthor struct {
	Uid  uint64 `json:"uid"`            // 用户id
	Name string `json:"name"`           // 用户昵称
	Icon string `json:"icon,omitempty"` // 用户头像
}

type PostDataCategory struct {
	Id   uint64 `json:"id"`   // 标签id
	Name string `json:"name"` // 标签名称
}

package modelv

type PostListPistageReq struct {
	// Uid      uint64 `form:"uid"`
	Page     uint64 `form:"page"`
	PostType uint64 `form:"postType"`
	// Category uint64 `form:"category"` // 分类筛选
}

func (req *PostListPistageReq) Init() {
	if req == nil {
		req = &PostListPistageReq{}
	}
	if req.Page == 0 {
		req.Page = 1
	}
}

// func (req *PostReq) Check()

type PostListPageResp struct {
	IsLastPage bool        `json:"isLastPage"` // 是否最后一页
	PostList   []*PostData `json:"postList"`   // 帖子列表
}

type PostData struct {
	PostId      uint64              `json:"postId"`             // 帖子id
	PostType    uint64              `json:"postType"`           // 帖子类型
	PostTitle   string              `json:"postTitle"`          // 标题
	PostContent string              `json:"postContent"`        // 内容
	Media       []*PostDataMedia    `json:"media,omitempty"`    // 图片、视频
	Author      *PostDataAuthor     `json:"author"`             // 作者
	LikeCnt     uint64              `json:"LikeCnt"`            // 点赞
	ReplyCnt    uint64              `json:"ReplyCnt"`           // 回复
	CollectCnt  uint64              `json:"collectCnt"`         // 收藏
	CreateTime  uint64              `json:"createTime"`         // 时间
	Category    []*PostDataCategory `json:"category"`           // 标签
	OnDoor      int                 `json:"onDoror,omitempty"`  // 是否可上门
	Price       float64             `json:"price,omitempty"`    // 价格
	Location    string              `json:"location,omitempty"` // 地址
}

type PostDataMedia struct {
	MType    int    `json:"mType"`    // 1:图片；2:视频
	PicUrl   string `json:"picUrl"`   // 如果是视频，该字段为视频封面地址
	VideoUrl string `json:"videoUrl"` // 视频地址
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

type PostCreateReq struct {
	Uid     int    `form:"-"` // 用户id
	UToken  string `form:"-"` // 登录token
	Sign    string `form:"-"` // 签名
	TimeStr string `form:"-"` // 客户端请求时间
}

type PostCreateBody struct {
	PostType int              `json:"postType"` // 帖子类型
	Title    string           `json:"title"`    // 标题
	Content  string           `json:"content"`  // 内容
	Media    []*PostDataMedia `json:"media"`    // 图片/视频
	Category []int            `json:"category"` // 标签id列表
	OnDoor   int              `json:"onDoror"`  // 0 : 不可上门，1：可以
	Price    float64          `json:"price"`    // 价格
	Location string           `json:"location"` // 地址
}

func (b *PostCreateBody) Check(postType int) bool {
	if b.Title == "" || b.Content == "" || len(b.Category) == 0 {
		return false
	}
	return true
}

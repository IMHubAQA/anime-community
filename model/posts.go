package model

const (
	ANIMEPOST_STATUS_VALID = iota + 1 // 正常
	ANIMEPOST_STATUS_DELET            // 删除
)

const (
	ANIMEPOST_TYPE_NORMAL = iota + 1 // 普通贴
	ANIMEPOST_TYPE_MAKEUP            // 约妆
)

type AnimePost struct {
	Id         int64  `json:"id" gorm:"column:id"`                                        //
	PostType   uint8  `json:"postType" gorm:"column:post_type"`                           // 帖子类型
	UserId     int64  `json:"user_id" gorm:"column:user_id"`                              // 用户id
	Title      string `json:"title" gorm:"column:title"`                                  // 标题
	Content    string `json:"content" gorm:"column:content"`                              // 内容
	Media      string `json:"media" gorm:"column:media"`                                  // 图片/视频
	AppendInfo string `json:"append_info" gorm:"column:append_info"`                      // 附加信息
	Status     uint8  `json:"status" gorm:"column:status"`                                //
	CreateTime int64  `json:"create_time" gorm:"autoCreateTime:milli;column:create_time"` //
	UpdateTime int64  `json:"update_time" gorm:"autoUpdateTime:milli;column:update_time"` //
}

func (a AnimePost) TableName() string {
	return "anime_post"
}

type AnimePostCategory struct {
	Id         int64  `json:"id" gorm:"column:id"`                                        //
	PostType   uint8  `json:"post_type" gorm:"column:post_type"`                          //
	Name       string `json:"name" gorm:"column:name"`                                    //
	Status     uint8  `json:"status" gorm:"column:status"`                                //
	CreateTime int64  `json:"create_time" gorm:"autoCreateTime:milli;column:create_time"` //
	UpdateTime int64  `json:"update_time" gorm:"autoUpdateTime:milli;column:update_time"` //
}

func (a AnimePostCategory) TableName() string {
	return "anime_post_category"
}

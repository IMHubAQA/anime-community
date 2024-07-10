package modele

const (
	ANIMECOMMENT_REPLYTYPE_POST = iota + 1
	ANIMECOMMENT_REPLYTYPE_COMMENT
)

const (
	ANIMECOMMENT_STATUS_VALID = iota + 1 // 正常
	ANIMECOMMENT_STATUS_DELET            // 删除
)

type AnimeComment struct {
	Id           int64  `json:"id" gorm:"column:id"`                         //
	Content      string `json:"content" gorm:"column:content"`               // 内容
	UserId       int64  `json:"user_id" gorm:"column:user_id"`               // 回复者id
	PostId       int64  `json:"post_id" gorm:"column:post_id"`               // 帖子id
	ReplyType    uint8  `json:"reply_type" gorm:"column:reply_type"`         // 回复类型
	ReplyId      int64  `json:"reply_id" gorm:"column:reply_id"`             // 回复类型id
	TargetUserId int64  `json:"target_user_id" gorm:"column:target_user_id"` // 回复用户id
	Status       uint8  `json:"status" gorm:"column:status"`                 //
	CreateTime   int64  `json:"create_time" gorm:"column:create_time"`       //
}

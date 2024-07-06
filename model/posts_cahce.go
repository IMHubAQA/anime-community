package model

// 标签缓存
type PostCategoryCache struct {
	Id       int64  `json:"id"`
	PostType uint8  `json:"postType"`
	Name     string `json:"name"`
}

func NewPostCategoryCache(entity *AnimePostCategory) *PostCategoryCache {
	if entity == nil {
		return nil
	}
	return &PostCategoryCache{
		Id:       entity.Id,
		PostType: entity.PostType,
		Name:     entity.Name,
	}
}

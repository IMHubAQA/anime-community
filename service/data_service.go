package service

import (
	modelc "anime-community/model/cache"
)

type DataEngine struct {
	UserInfo     interface{}
	CategoryInfo *modelc.PostCategoryCache
}

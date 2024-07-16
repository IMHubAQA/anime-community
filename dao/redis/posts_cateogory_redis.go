package redis

import (
	"context"
	"encoding/json"
	"fmt"

	modelc "anime-community/model/cache"
)

// 保存帖子标签信息
func SetPostsCategory(ctx context.Context, postType int, cacheValues []*modelc.PostCategoryCache) error {
	if len(cacheValues) == 0 {
		return nil
	}
	keys, values := make([]string, len(cacheValues)+1), make([]string, len(cacheValues)+1)

	for i, cacheValue := range cacheValues {
		keys[i] = PostCategoryRedisKey.GetKey(cacheValue.Id)
		b, _ := json.Marshal(cacheValue)
		values[i] = string(b)
	}

	keys[len(keys)-1] = PostCategoryRedisKey.GetKey("list", postType)
	b, _ := json.Marshal(cacheValues)
	values[len(values)-1] = string(b)

	return MutiSet(ctx, keys, values, PostCategoryRedisKey.GetExpire())
}

// 获取帖子标签信息
func GetPostsCategory(ctx context.Context, ids []int) (map[int]*modelc.PostCategoryCache, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("empty ids")
	}
	keys := make([]string, len(ids))
	keyMap := make(map[string]int)
	for i, id := range ids {
		key := PostCategoryRedisKey.GetKey(id)
		if _, ok := keyMap[key]; ok {
			continue
		}
		keys[i] = key
		keyMap[key] = id
	}

	m, err := MutiGet(ctx, keys)
	if err != nil {
		return nil, err
	}
	res := make(map[int]*modelc.PostCategoryCache)
	for key, value := range m {
		pcc := &modelc.PostCategoryCache{}
		err := json.Unmarshal([]byte(value), pcc)
		if err != nil {
			continue
		}
		res[keyMap[key]] = pcc
	}
	return res, nil
}

func GetPostsCategoryList(ctx context.Context, postType int) ([]*modelc.PostCategoryCache, error) {
	b, err := GetCommunityClient().Get(ctx, PostCategoryRedisKey.GetKey("list", postType)).Result()
	if err != nil {
		return nil, err
	}
	res := []*modelc.PostCategoryCache{}
	err = json.Unmarshal([]byte(b), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

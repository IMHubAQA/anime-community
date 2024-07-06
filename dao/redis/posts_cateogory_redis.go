package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bytedance/sonic"

	"anime-community/model"
)

// 保存帖子标签信息
func SetPostsCategory(ctx context.Context, cacheValues []*model.PostCategoryCache) error {
	if len(cacheValues) == 0 {
		return nil
	}

	keys, values := make([]string, 0, len(cacheValues)), make([]string, 0, len(cacheValues))
	for i, cacheValue := range cacheValues {
		keys[i] = PostCategoryRedisKey.GetKey(strconv.Itoa(int(cacheValue.Id)))
		b, _ := sonic.Marshal(cacheValue)
		values[i] = string(b)
	}

	return MutiSet(ctx, keys, values, PostCategoryRedisKey.GetExpire())
}

// 获取帖子标签信息
func GetPostsCategory(ctx context.Context, ids []int) (map[int]*model.PostCategoryCache, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("empty ids")
	}
	keys := make([]string, 0, len(ids))
	keyMap := make(map[string]int)
	for i, id := range ids {
		key := PostCategoryRedisKey.GetKey(strconv.Itoa(int(id)))
		keys[i] = key
		keyMap[key] = id
	}

	m, err := MutiGet(ctx, keys)
	if err != nil {
		return nil, err
	}
	res := make(map[int]*model.PostCategoryCache)
	for key, value := range m {
		pcc := &model.PostCategoryCache{}
		err := sonic.Unmarshal([]byte(value), pcc)
		if err != nil {
			continue
		}
		res[keyMap[key]] = pcc
	}
	return res, nil
}

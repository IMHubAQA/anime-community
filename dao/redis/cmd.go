package redis

import (
	"context"
	"fmt"
)

const (
	_SINGLE_MUTI_SIZE = 20
)

func splitKeys(curr int, length int) int {
	if curr+_SINGLE_MUTI_SIZE >= length {
		return length
	}
	return curr + _SINGLE_MUTI_SIZE
}

func MutiGet(ctx context.Context, keys []string) (map[string]string, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("empty key")
	}
	offset := 0
	length := len(keys)
	m := make(map[string]string)
	for offset < length {
		offset = splitKeys(offset, length)
	}
	return m, nil
}

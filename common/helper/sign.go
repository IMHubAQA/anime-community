package helper

import (
	"context"
	"encoding/hex"
	"hash"

	"anime-community/common/logs"
)

func CheckSign(ctx context.Context, sign string, hashFunc hash.Hash, params ...string) bool {
	for _, param := range params {
		hashFunc.Write([]byte(param))
	}
	validSign := hex.EncodeToString(hashFunc.Sum(nil))
	if sign != validSign {
		logs.Infof(ctx, "CheckSign fail. sign=%v validSign=%v params=%v", sign, validSign, params)
	}
	return sign == validSign
}

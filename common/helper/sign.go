package helper

import (
	"encoding/hex"
	"hash"
)

func CheckSign(sign string, hashFunc hash.Hash, params ...string) bool {
	for _, param := range params {
		hashFunc.Write([]byte(param))
	}
	return sign == hex.EncodeToString(hashFunc.Sum(nil))
}

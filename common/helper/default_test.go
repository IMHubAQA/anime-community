package helper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGenSignByCommon(t *testing.T) {
	uid := "1"
	timeStr := "1"
	jsonBody := `{"postId":19,"relayType":2,"relayId":2,"Content":"评论！@#￥￥%~~dsada~"}`
	h := sha256.New()
	h.Write([]byte(uid))
	h.Write([]byte(timeStr))
	h.Write([]byte(jsonBody))
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign, CheckSign(context.Background(), sign, sha256.New(), uid, timeStr, jsonBody))
}

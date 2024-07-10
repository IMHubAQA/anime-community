package modelv

import (
	"context"
	"crypto/sha256"
	"strconv"

	beegoctx "github.com/beego/beego/v2/server/web/context"

	"anime-community/common/constants"
	"anime-community/common/helper"
)

type BaseHeader struct {
	Uid     int    `form:"-"` // 用户id
	UToken  string `form:"-"` // 登录token
	Sign    string `form:"-"` // 签名
	TimeStr string `form:"-"` // 客户端请求时间
}

func GetAndCheckBaseHeader(ctx context.Context, beegoCtx *beegoctx.Context) (*BaseHeader, *constants.Error) {
	header := &BaseHeader{
		UToken:  beegoCtx.Request.Header.Get("uToken"),
		Sign:    beegoCtx.Request.Header.Get("sign"),
		TimeStr: beegoCtx.Request.Header.Get("timeStr"),
	}
	uid := beegoCtx.Request.Header.Get("uid")
	header.Uid, _ = strconv.Atoi(uid)

	if header.Uid <= 0 || header.UToken == "" {
		return nil, constants.InvalidParamsError

	}

	if !helper.CheckSign(header.Sign, sha256.New(), uid, header.TimeStr, string(beegoCtx.Input.RequestBody)) {
		return nil, constants.InvalidSignError
	}
	return header, nil
}

type MediaData struct {
	MType    int    `json:"mType"`    // 1:图片；2:视频
	PicUrl   string `json:"picUrl"`   // 如果是视频，该字段为视频封面地址
	VideoUrl string `json:"videoUrl"` // 视频地址
}

type AuthorData struct {
	Uid  uint64 `json:"uid"`            // 用户id
	Name string `json:"name"`           // 用户昵称
	Icon string `json:"icon,omitempty"` // 用户头像
}

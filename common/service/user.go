package commservice

import (
	"anime-community/common/httpc"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const (
	_USER_MGETUSERINFO_URL = "http://122.51.70.205:8000/api/getUsers"
	__USER_VERFIYTOKEN_URL = "http://122.51.70.205:8000/api/verifyToken"
)

type UserData struct {
	ID        int         `json:"id"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Salt      string      `json:"salt"`
	NickName  string      `json:"nickName"`
	AvatarURL string      `json:"avatarUrl"`
	Gender    int         `json:"gender"`
	BirthDate string      `json:"birthDate"`
}

type MGetUserReq struct {
	UserIds []int64 `json:"userIds"`
}

type MGetUserResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data []*UserData `json:"data"`
}

func MGetUserInfos(ctx context.Context, req *MGetUserReq, timeout time.Duration) ([]*UserData, error) {
	data, err := httpc.PostJson(ctx, _USER_MGETUSERINFO_URL, nil, req, timeout)
	if err != nil {
		return nil, err
	}
	resp := &MGetUserResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf("code=%v, msg=%v", resp.Code, resp.Msg)
	}
	return resp.Data, nil
}

func GetUserInfo(ctx context.Context, uid int64, timeout time.Duration) (*UserData, error) {
	req := &MGetUserReq{UserIds: []int64{uid}}
	datas, err := MGetUserInfos(ctx, req, timeout)
	if err != nil {
		return nil, err
	}
	if len(datas) == 0 {
		return nil, fmt.Errorf("uid not exist id=%v", uid)
	}
	return datas[0], nil
}

func MGetUserInfom(ctx context.Context, req *MGetUserReq, timeout time.Duration) (map[int]*UserData, error) {
	datas, err := MGetUserInfos(ctx, req, timeout)
	if err != nil {
		return nil, err
	}
	resp := make(map[int]*UserData)
	for _, data := range datas {
		resp[data.ID] = data
	}
	return resp, nil
}

type VerifyTokenReq struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}

type VerifyTokenResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func VerifyToken(ctx context.Context, req *VerifyTokenReq, timeout time.Duration) (bool, error) {
	data, err := httpc.PostJson(ctx, __USER_VERFIYTOKEN_URL, nil, req, timeout)
	if err != nil {
		return false, err
	}
	resp := &MGetUserResp{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return false, err
	}
	if resp.Code != 200 {
		return false, fmt.Errorf("code=%v, msg=%v", resp.Code, resp.Msg)
	}
	return true, nil
}

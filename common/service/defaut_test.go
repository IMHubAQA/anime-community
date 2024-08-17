package commservice

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMgetUserInfo(t *testing.T) {
	req := &MGetUserReq{
		UserIds: []int64{1},
	}
	data, err := MGetUserInfos(context.Background(), req, time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v,%v \n", data[0], err)
}

func TestVerifyToken(t *testing.T) {
	req := &VerifyTokenReq{
		UserId: 1,
		Token:  "token",
	}
	data, err := VerifyToken(context.Background(), req, time.Second)
	fmt.Printf("%+v,%v \n", data, err)
}

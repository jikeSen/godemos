package bfPost

import (
	"github.com/jikeSen/godemos/MyCurl"
	"github.com/jikeSen/godemos/router"
)

type GetChookParam struct {
	Uid   int `json:"uid"`
	Level int `json:"level"`
	Ltd   int `json:"Ltd"`
}

// 抢购金鸡
func GetChook(uid int) string {
	params := GetChookParam{
		uid,
		1,
		uid,
	}

	urls := router.GetChook

	return MyCurl.HttpPost(urls, params)
}

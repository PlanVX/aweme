package types

import (
	"io"
)

// FeedReq is the request of feed api
type FeedReq struct {
	LatestTime int64  `form:"latest_time" json:"latest_time,omitempty" query:"latest_time"` //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `form:"token" json:"token,omitempty" query:"token"`                   // 登录状态则有
}

// FeedResp is the response of feed api
type FeedResp struct {
	Response
	NextTime  int64    `json:"next_time"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []*Video `json:"video_list"` // 视频列表
}

// UploadReq is the request of upload api /publish/action/
type UploadReq struct {
	Title    string `form:"title" json:"title"`
	FileName string
	Data     io.Reader
}

// UploadResp is the response of upload api /publish/action/
type UploadResp struct {
	Response
}

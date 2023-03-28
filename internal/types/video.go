/*
 * Copyright (c) 2023 The PlanVX Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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

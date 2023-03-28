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

// FavoriteActionReq 用户点赞请求
type FavoriteActionReq struct {
	Token      string `query:"token" json:"token" form:"token" mapstructure:"token"`                         // 用户鉴权token
	VideoID    int64  `query:"video_id" json:"video_id" form:"video_id" mapstructure:"video_id"`             // 视频id
	ActionType int32  `query:"action_type" json:"action_type" form:"action_type" mapstructure:"action_type"` // 1-点赞，2-取消点赞
}

// FavoriteActionResp 用户点赞响应
type FavoriteActionResp struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

// FavoriteListReq 用户点赞列表请求
type FavoriteListReq struct {
	UserID int64  `query:"user_id"` // 用户id
	Token  string `query:"token"`   // 用户鉴权token
}

// FavoriteListResp 用户点赞列表响应
type FavoriteListResp struct {
	StatusCode int32    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `json:"status_msg"`  // 返回状态描述
	VideoList  []*Video `json:"video_list"`  // 用户点赞视频列表
}

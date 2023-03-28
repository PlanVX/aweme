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

// PublishListReq is the request of publish list api
type PublishListReq struct {
	Token  string `query:"token" form:"token" json:"token"`       // token
	UserID int64  `query:"user_id" form:"user_id" json:"user_id"` // user id
}

// PublishListResp is the response of publish list api
type PublishListResp struct {
	Response
	VideoList []*Video `json:"video_list"` // 用户发布的视频列表
}

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

// UserResp is the response of user register/login api
type UserResp struct {
	Response
	Token  string `json:"token"`   // 用户鉴权token
	UserID int64  `json:"user_id"` // 用户id
}

// UserReq is the request of user register/login api
type UserReq struct {
	Username string `form:"username" json:"username" query:"username" validate:"required,printascii,min=1,max=16"` // 用户名
	Password string `form:"password" json:"password" query:"password" validate:"required,printascii,min=6,max=16"` // 密码
}

// UserInfoReq is the request of user info api
type UserInfoReq struct {
	UserID int64  `query:"user_id" binding:"required"`
	Token  string `query:"token" binding:"required"`
}

// UserInfoResp is the response of user info api
type UserInfoResp struct {
	Response
	User *User `json:"user,omitempty"` // 用户信息
}

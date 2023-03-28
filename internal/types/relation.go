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

// RelationFollowListResp 用于获取关注列表响应
type RelationFollowListResp struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	UserList   []*User `json:"user_list"`   // 用户信息列表
}

// RelationFollowerListResp 用于获取粉丝列表响应
type RelationFollowerListResp struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	UserList   []*User `json:"user_list"`   // 用户信息列表
}

// RelationFriendListReq 用于获取好友列表请求
type RelationFriendListReq struct {
	UserID int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// RelationFriendListResp 用于获取好友列表响应
type RelationFriendListResp struct {
	StatusCode int32            `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string           `json:"status_msg"`  // 返回状态描述
	UserList   []FriendUserResp `json:"user_list"`   // 用户信息列表
}

// FriendUserResp 好友用户信息
type FriendUserResp struct {
	User
	Message string `json:"message,omitempty"` // 和该好友的最新聊天消息
	MsgType int64  `json:"msg_type"`          // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

// RelationActionReq 关系操作请求
type RelationActionReq struct {
	Token      string `json:"token"`       // 用户鉴权 token
	ToUserID   int64  `json:"to_user_id"`  // 对方用户 id
	ActionType int32  `json:"action_type"` // 1-关注，2-取消关注
}

// RelationActionResp 关系操作响应
type RelationActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// RelationFollowListReq 用于获取关注列表请求
type RelationFollowListReq struct {
	UserID int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// RelationFollowerListReq 用于获取粉丝列表请求
type RelationFollowerListReq struct {
	UserID int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

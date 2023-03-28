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

// MessageListReq chat message list request
type MessageListReq struct {
	Token    string `json:"token" query:"token"`           // 用户鉴权token
	ToUserID int64  `json:"to_user_id" query:"to_user_id"` // 对方用户id
}

// MessageListResp chat message list response
type MessageListResp struct {
	StatusCode  int32      `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string     `json:"status_msg"`   // 返回状态描述
	MessageList []*Message `json:"message_list"` // 消息列表
}

// MessageActionReq chat action request
type MessageActionReq struct {
	Token      string `json:"token" query:"token"`             // 用户鉴权token
	ToUserID   int64  `json:"to_user_id" query:"to_user_id"`   // 对方用户id
	ActionType int32  `json:"action_type" query:"action_type"` // 1-发送消息
	Content    string `json:"content" query:"content"`         // 消息内容
}

// MessageActionResp chat action response
type MessageActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

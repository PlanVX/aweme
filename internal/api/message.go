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

package api

import (
	"context"
	"github.com/PlanVX/aweme/internal/types"
)

// NewMessageChat godoc
// @Summary 聊天记录
// @Description 聊天记录
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param message query types.MessageListReq true "用户信息"
// @Success 200 {object} types.MessageListResp
// @Router /message/chat/ [get]
func NewMessageChat() *API {
	return &API{
		Method: "GET",
		Path:   "/message/chat/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.MessageListReq) (*types.MessageListResp, error) {
			return nil, nil
		}),
	}
}

// NewMessageAction godoc
// @Summary 消息操作
// @Description 消息操作
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param message query types.MessageActionReq true "用户消息"
// @Success 200 {object} types.MessageActionResp
// @Router /message/action/ [post]
func NewMessageAction() *API {
	return &API{
		Method: "POST",
		Path:   "/message/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.MessageActionReq) (*types.MessageActionResp, error) {
			return nil, nil
		}),
	}
}

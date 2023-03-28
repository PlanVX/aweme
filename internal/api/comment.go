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
	"github.com/PlanVX/aweme/internal/logic"
)

// NewCommentAction godoc
// @Summary 评论操作
// @Description 评论操作
// @Tags 互动接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user formData types.CommentActionReq true "评论信息"
// @Success 200 {object} types.CommentActionResp
// @Router /comment/action/ [post]
func NewCommentAction(action *logic.CommentAction) *API {
	return &API{
		Method:  "POST",
		Path:    "/comment/action/",
		Handler: WrapperFunc(action.CommentAction),
	}
}

// NewCommentList godoc
// @Summary 评论列表
// @Description 评论列表
// @Tags 互动接口
// @Produce json
// @Param user_id query types.CommentListReq true "用户信息"
// @Success 200 {object} types.CommentListResp
// @Router /comment/list/ [get]
func NewCommentList(list *logic.CommentList) *API {
	return &API{
		Method:  "GET",
		Path:    "/comment/list/",
		Handler: WrapperFunc(list.CommentList),
	}
}

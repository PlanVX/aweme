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

// NewFavoriteAction godoc
// @Summary 赞操作
// @Description 赞操作
// @Tags 互动接口
// @Produce json
// @Param favorite formData types.FavoriteActionReq true "用户消息信息"
// @Success 200 {object} types.FavoriteActionResp
// @Router /favorite/action/ [POST]
func NewFavoriteAction(like *logic.Like) *API {
	return &API{
		Method:  "POST",
		Path:    "/favorite/action/",
		Handler: WrapperFunc(like.Like),
	}
}

// NewFavoriteList godoc
// @Summary 点赞列表
// @Description 点赞列表
// @Tags 互动接口
// @Produce json
// @Param favorite query types.FavoriteListReq true "请求信息"
// @Success 200 {object} types.FavoriteListResp
// @Router /favorite/list/ [get]
func NewFavoriteList(list *logic.LikeList) *API {
	return &API{
		Method:  "GET",
		Path:    "/favorite/list/",
		Handler: WrapperFunc(list.LikeList),
	}
}

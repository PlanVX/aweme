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

// NewPublishList godoc
// @Summary 获取视频列表
// @Description 获取视频列表
// @Tags 基础接口
// @Produce json
// @Param param query types.PublishListReq true "获取视频列表参数"
// @Success 200 {object} types.PublishListResp
// @Router /publish/list/ [get]
func NewPublishList(list *logic.PublishList) *API {
	return &API{
		Method:  "GET",
		Path:    "/publish/list/",
		Handler: WrapperFunc(list.PublishList),
	}
}

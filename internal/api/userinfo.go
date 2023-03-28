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

import "github.com/PlanVX/aweme/internal/logic"

// NewUserInfo godoc
// @Summary 用户信息
// @Description 用户信息
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param user query types.UserInfoReq true "用户信息"
// @Success 200 {object} types.UserInfoResp
// @Router /user/ [get]
func NewUserInfo(profile *logic.UserProfile) *API {
	return &API{
		Path:    "/user/",
		Method:  "GET",
		Handler: WrapperFunc(profile.GetProfile),
	}
}

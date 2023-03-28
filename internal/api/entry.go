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

// NewRegister godoc
// @Summary 用户注册
// @Description 用户注册
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param user formData types.UserReq true "用户信息"
// @Success 200 {object} types.UserResp
// @Router /user/register/ [post]
func NewRegister(l *logic.Register) *API {
	return &API{
		Method:  "POST",
		Path:    "/user/register/",
		Handler: WrapperFunc(l.Register),
	}
}

// NewLogin godoc
// @Summary 用户登陆
// @Description 用户登陆
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param user formData types.UserReq true "用户信息"
// @Success 200 {object} types.UserResp
// @Router /user/login/ [post]
func NewLogin(l *logic.Login) *API {
	return &API{
		Method:  "POST",
		Path:    "/user/login/",
		Handler: WrapperFunc(l.Login),
	}
}

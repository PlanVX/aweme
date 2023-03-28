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

package logic

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"go.uber.org/fx"
)

type (
	// UserProfile is the logic for querying user profile
	UserProfile struct {
		userModel dal.UserQuery
	}
	// UserProfileParam is the parameter for NewUserProfile
	UserProfileParam struct {
		fx.In
		UserModel dal.UserQuery
	}
)

// NewUserProfile returns a new UserProfile logic
func NewUserProfile(param UserProfileParam) *UserProfile {
	return &UserProfile{
		userModel: param.UserModel,
	}
}

// GetProfile 获取用户信息
// 根据用户 id 获取用户信息
func (u *UserProfile) GetProfile(ctx context.Context, req *types.UserInfoReq) (*types.UserInfoResp, error) {

	v, err := u.userModel.FindOne(ctx, req.UserID) // 根据用户 id 获取用户信息
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		User: covertUser(v),
	}, nil
}

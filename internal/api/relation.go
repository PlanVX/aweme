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

// NewRelationAction godoc
// @Summary 关系操作
// @Description 关系操作
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData types.RelationActionReq true "用户信息"
// @Success 200 {object} types.RelationActionResp
// @Router /relation/action/ [post]
func NewRelationAction() *API {
	return &API{
		Method: "POST",
		Path:   "/relation/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.RelationActionReq) (*types.RelationActionResp, error) {
			return nil, nil
		}),
	}
}

// NewRelationFollowList godoc
// @Summary 用户关注列表
// @Description 用户关注列表
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData types.RelationFollowListReq true "用户信息"
// @Success 200 {object} types.RelationFollowListResp
// @Router /relation/follow/list/ [get]
func NewRelationFollowList() *API {
	return &API{
		Method: "GET",
		Path:   "/relation/follow/list//",
		Handler: WrapperFunc(func(ctx context.Context, req *types.RelationFollowListReq) (*types.RelationFollowListResp, error) {
			return nil, nil
		}),
	}
}

// NewRelationFollowerList godoc
// @Summary 用户粉丝列表
// @Description 用户粉丝列表
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData types.RelationFollowerListReq true "用户信息"
// @Success 200 {object} types.RelationFollowerListResp
// @Router /relation/follower/list/ [get]
func NewRelationFollowerList() *API {
	return &API{
		Method: "GRT",
		Path:   "/relation/follower/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.RelationFollowerListReq) (*types.RelationFollowerListResp, error) {
			return nil, nil
		}),
	}
}

// NewRelationFriendList godoc
// @Summary 用户好友列表
// @Description 用户好友列表
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData types.RelationFriendListReq true "用户信息"
// @Success 200 {object} types.RelationFriendListResp
// @Router /relation/friend/list/ [get]
func NewRelationFriendList() *API {
	return &API{
		Method: "GET",
		Path:   "/relation/friend/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.RelationFriendListReq) (*types.RelationFriendListResp, error) {
			return nil, nil
		}),
	}
}

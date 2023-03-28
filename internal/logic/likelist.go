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
	// LikeList is the like list logic layer struct
	LikeList struct {
		likeModel  dal.LikeQuery
		videoModel dal.VideoQuery
		userModel  dal.UserQuery
	}
	// LikeListParam is the param for NewLikeList
	LikeListParam struct {
		fx.In
		LikeModel  dal.LikeQuery
		UserModel  dal.UserQuery
		VideoModel dal.VideoQuery
	}
)

// NewLikeList returns a new LikeList logic
func NewLikeList(param LikeListParam) *LikeList {
	return &LikeList{
		likeModel:  param.LikeModel,
		userModel:  param.UserModel,
		videoModel: param.VideoModel,
	}
}

// LikeList is the like list logic
// handle the like list
func (l *LikeList) LikeList(ctx context.Context, req *types.FavoriteListReq) (*types.FavoriteListResp, error) {
	owner, _ := ctx.Value(ContextKey).(int64)

	likes, err := l.likeModel.FindVideoIDsByUserID(ctx, req.UserID, 30, 0)
	if err != nil {
		return nil, err
	}

	many, err := l.videoModel.FindMany(ctx, likes)
	if err != nil {
		return nil, err
	}

	users, err := l.userModel.FindMany(ctx, extractUserIDs(many)) // 根据用户 id 批量查询用户信息
	if err != nil {
		return nil, err
	}

	var likedList []int64
	if owner != 0 {
		likedList, err = l.likeModel.FindWhetherLiked(ctx, owner, likes)
		if err != nil {
			return nil, err
		}
	}
	// set is favorite

	videos := packVideos(many, users, likedList)

	return &types.FavoriteListResp{
		VideoList: videos,
	}, nil
}

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
	// PublishList is the logic for publish list
	PublishList struct {
		videoQuery dal.VideoQuery
		userQuery  dal.UserQuery
		likeQuery  dal.LikeQuery
	}

	// PublishListParam is the parameter for NewPublishList
	PublishListParam struct {
		fx.In
		VideoQuery dal.VideoQuery
		UserQuery  dal.UserQuery
		LikeQuery  dal.LikeQuery
	}
)

// NewPublishList returns a new PublishList logic
func NewPublishList(param PublishListParam) *PublishList {
	return &PublishList{
		videoQuery: param.VideoQuery,
		userQuery:  param.UserQuery,
		likeQuery:  param.LikeQuery,
	}
}

// PublishList gets publish list
func (p *PublishList) PublishList(ctx context.Context, req *types.PublishListReq) (*types.PublishListResp, error) {

	owner, _ := ctx.Value(ContextKey).(int64)

	// query videos by specified user id
	filteredVideos, err := p.videoQuery.FindByUserID(ctx, req.UserID, 30, 0)
	if err != nil {
		return nil, err
	}

	// query user info by specified user id
	user, err := p.userQuery.FindOne(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	// video id list
	videoIDs := extractVideosIDs(filteredVideos)

	// query like info by specified user id
	var likes []int64
	if owner != 0 {
		likes, err = p.likeQuery.FindWhetherLiked(ctx, owner, videoIDs)
		if err != nil {
			return nil, err
		}
	}

	videos := packVideos(filteredVideos, []*dal.User{user}, likes)

	return &types.PublishListResp{VideoList: videos}, nil
}

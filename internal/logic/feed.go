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
	// Feed is the logic for feed
	Feed struct {
		videoQuery dal.VideoQuery
		userQuery  dal.UserQuery
		likeQuery  dal.LikeQuery
	}
	// FeedParam is the parameter for NewFeed
	FeedParam struct {
		fx.In
		VideoQuery dal.VideoQuery
		UserQuery  dal.UserQuery
		LikeQuery  dal.LikeQuery
	}
)

// NewFeed creates a new feed logic
func NewFeed(param FeedParam) *Feed {
	return &Feed{
		videoQuery: param.VideoQuery,
		userQuery:  param.UserQuery,
		likeQuery:  param.LikeQuery,
	}
}

// Feed 获取首页视频流
func (f *Feed) Feed(ctx context.Context, req *types.FeedReq) (resp *types.FeedResp, err error) {
	owner, _ := ctx.Value(ContextKey).(int64)

	latestVideo, err := f.videoQuery.FindLatest(ctx, req.LatestTime, 30) // 查询 30 个视频
	if err != nil {
		return nil, err
	}

	users, err := f.userQuery.FindMany(ctx, extractUserIDs(latestVideo)) // 根据用户 id 批量查询用户信息
	if err != nil {
		return nil, err
	}

	videoIDs := extractVideosIDs(latestVideo)

	var likedList []int64
	if owner != 0 {
		likedList, err = f.likeQuery.FindWhetherLiked(ctx, owner, videoIDs)
		if err != nil {
			return nil, err
		}
	}

	videos := packVideos(latestVideo, users, likedList)

	return &types.FeedResp{
		VideoList: videos,
	}, nil
}

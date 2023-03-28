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
	"errors"
	"testing"

	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/PlanVX/aweme/internal/dal"
)

func TestLikeList(t *testing.T) {
	assertions := assert.New(t)
	videoIDs := []int64{1, 2}
	dalVideos := []*dal.Video{{ID: 1, UserID: 4}, {ID: 2, UserID: 5}}
	dalUsers := []*dal.User{{ID: 4}, {ID: 5}}
	likedList := []int64{1}
	ctx := ContextWithOwner(1)
	t.Run("query success", func(t *testing.T) {
		likeModel, videoModel, userModel, list := mockLikeList(t)
		likeModel.On("FindVideoIDsByUserID",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(videoIDs, nil)
		likeModel.On("FindWhetherLiked", mock.Anything, mock.Anything, mock.Anything).Return(likedList, nil)
		videoModel.On("FindMany", mock.Anything, mock.Anything).Return(dalVideos, nil)
		userModel.On("FindMany", mock.Anything, mock.Anything).Return(dalUsers, nil)

		resp, err := list.LikeList(ctx, &types.FavoriteListReq{
			UserID: 1,
		})
		assertions.NoError(err)
		assertions.NotNil(resp.VideoList)
		assertions.True(resp.VideoList[0].IsFavorite)
		assertions.False(resp.VideoList[1].IsFavorite)
		for i, item := range resp.VideoList {
			assertions.Equal(dalVideos[i].ID, item.ID)
			assertions.Equal(dalVideos[i].UserID, item.Author.ID)
		}
	})
	t.Run("query like list failed", func(t *testing.T) {
		likeModel, _, _, list := mockLikeList(t)
		likeModel.On("FindVideoIDsByUserID",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(nil, errors.New("query like list failed"))

		_, err := list.LikeList(ctx, &types.FavoriteListReq{
			UserID: 1,
		})
		assertions.Error(err)
	})
	t.Run("query video list failed", func(t *testing.T) {
		likeModel, videoModel, _, list := mockLikeList(t)
		likeModel.On("FindVideoIDsByUserID",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(videoIDs, nil)
		videoModel.On("FindMany", mock.Anything, mock.Anything).Return(nil, errors.New("query video list failed"))

		_, err := list.LikeList(ctx, &types.FavoriteListReq{
			UserID: 1,
		})
		assertions.Error(err)
	})
	t.Run("query user list failed", func(t *testing.T) {
		likeModel, videoModel, userModel, list := mockLikeList(t)
		likeModel.On("FindVideoIDsByUserID",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(videoIDs, nil)
		videoModel.On("FindMany", mock.Anything, mock.Anything).Return(dalVideos, nil)
		userModel.On("FindMany", mock.Anything, mock.Anything).Return(nil, errors.New("query user list failed"))

		_, err := list.LikeList(ctx, &types.FavoriteListReq{
			UserID: 1,
		})
		assertions.Error(err)
	})
}

func ContextWithOwner(owner int64) context.Context {
	ctx := context.WithValue(context.Background(), ContextKey, owner)
	return ctx
}

func mockLikeList(t *testing.T) (*LikeQuery, *VideoQuery, *UserQuery, *LikeList) {
	likeModel := NewLikeQuery(t)
	videoModel := NewVideoQuery(t)
	userModel := NewUserQuery(t)
	list := NewLikeList(LikeListParam{
		LikeModel:  likeModel,
		UserModel:  userModel,
		VideoModel: videoModel,
	})
	return likeModel, videoModel, userModel, list
}

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
	"errors"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPublishList(t *testing.T) {
	assertions := assert.New(t)
	videos := []*dal.Video{{ID: 1, UserID: 1}, {ID: 2, UserID: 1}, {ID: 3, UserID: 1}}
	user := &dal.User{ID: 1}
	req := &types.PublishListReq{UserID: user.ID}
	liked := []int64{1, 2}
	ctx := ContextWithOwner(int64(1))
	t.Run("Success", func(t *testing.T) {
		u, v, l, list := mockPublishList(t)
		v.On("FindByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(videos, nil)
		u.On("FindOne", mock.Anything, mock.Anything).Return(user, nil)
		l.On("FindWhetherLiked", mock.Anything, mock.Anything, mock.Anything).Return(liked, nil)
		resp, err := list.PublishList(ctx, req)
		assertions.NoError(err)
		assertions.Equal(3, len(resp.VideoList))
		assertions.Equal(videos[1].ID, resp.VideoList[1].ID)
		assertions.Equal(resp.VideoList[0].Author.ID, resp.VideoList[1].Author.ID)
		assertions.Equal(user.ID, resp.VideoList[0].Author.ID)
		assertions.Equal(true, resp.VideoList[0].IsFavorite)
		assertions.Equal(true, resp.VideoList[1].IsFavorite)
		assertions.Equal(false, resp.VideoList[2].IsFavorite)
	})
	t.Run("Video model error", func(t *testing.T) {
		_, v, _, list := mockPublishList(t)
		v.On("FindByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("video model error"))
		_, err := list.PublishList(ctx, req)
		assertions.Error(err)
	})
	t.Run("User model error", func(t *testing.T) {
		u, v, _, list := mockPublishList(t)
		v.On("FindByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(videos, nil)
		u.On("FindOne", mock.Anything, mock.Anything).Return(nil, errors.New("user model error"))
		_, err := list.PublishList(ctx, req)
		assertions.Error(err)
	})
	t.Run("Like model error", func(t *testing.T) {
		u, v, l, list := mockPublishList(t)
		v.On("FindByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(videos, nil)
		u.On("FindOne", mock.Anything, mock.Anything).Return(user, nil)
		l.On("FindWhetherLiked", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("like model error"))
		_, err := list.PublishList(ctx, req)
		assertions.Error(err)
	})
}

func mockPublishList(t *testing.T) (*UserQuery, *VideoQuery, *LikeQuery, *PublishList) {
	u, v, l := NewUserQuery(t), NewVideoQuery(t), NewLikeQuery(t)
	list := NewPublishList(PublishListParam{VideoQuery: v, UserQuery: u, LikeQuery: l})
	return u, v, l, list
}

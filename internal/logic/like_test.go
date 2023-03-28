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
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestLike(t *testing.T) {
	assertions := assert.New(t)
	t.Run("test add like success", func(t *testing.T) {
		model := NewLikeCommand(t)
		model.On("Insert", mock.Anything, mock.Anything).Return(nil)
		l := NewLike(LikeParam{LikeCommand: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 1, VideoID: 1})
		assertions.NoError(err)
		assertions.NotNil(resp)
	})
	t.Run("test add like failed", func(t *testing.T) {
		model := NewLikeCommand(t)
		model.On("Insert", mock.Anything, mock.Anything).Return(errors.New("failed"))
		l := NewLike(LikeParam{LikeCommand: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 1, VideoID: 1})
		assertions.Error(err)
		assertions.Nil(resp)
	})
	t.Run("test remove like success", func(t *testing.T) {
		model := NewLikeCommand(t)
		model.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		l := NewLike(LikeParam{LikeCommand: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 2, VideoID: 1})
		assertions.NoError(err)
		assertions.NotNil(resp)
	})
	t.Run("test remove like failed", func(t *testing.T) {
		model := NewLikeCommand(t)
		model.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed"))
		l := NewLike(LikeParam{LikeCommand: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 2, VideoID: 1})
		assertions.Error(err)
		assertions.Nil(resp)
	})
	t.Run("test invalid action type", func(t *testing.T) {
		model := NewLikeCommand(t)
		l := NewLike(LikeParam{LikeCommand: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 3, VideoID: 1})
		assertions.Error(err)
		assertions.Nil(resp)
	})
}

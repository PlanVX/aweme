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

func TestNewCommentAction(t *testing.T) {
	assertions := assert.New(t)
	var id int64 = 1
	ctx := ContextWithOwner(id)
	dalUser := &dal.User{ID: id}
	t.Run("create comment success", func(t *testing.T) {
		model := NewCommentCommand(t)
		user := NewUserQuery(t)
		model.On("Insert", mock.Anything, mock.Anything).Return(nil)
		user.On("FindOne", mock.Anything, mock.Anything).Return(dalUser, nil)
		c := NewCommentAction(CommentActionParam{CommentCommand: model, UserQuery: user})
		action, err := c.CommentAction(ctx, &types.CommentActionReq{
			VideoID:     1,
			ActionType:  1,
			CommentText: "hello",
		})
		assertions.NoError(err)
		assertions.NotNil(action)
		//action.Comment.CreateDate
		t.Logf("action.Comment.CreateDate: %s", action.Comment.CreateDate)
	})
	t.Run("create comment failed", func(t *testing.T) {
		model := NewCommentCommand(t)
		model.On("Insert", mock.Anything, mock.Anything).Return(errors.New("failed"))
		c := NewCommentAction(CommentActionParam{CommentCommand: model})
		action, err := c.CommentAction(ctx, &types.CommentActionReq{
			VideoID:     1,
			ActionType:  1,
			CommentText: "hello",
		})
		assertions.Error(err)
		assertions.Nil(action)
	})
	delReq := &types.CommentActionReq{
		ActionType: 2,
		CommentID:  1,
	}
	t.Run("delete comment success", func(t *testing.T) {
		model := NewCommentCommand(t)
		model.On("Delete", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
		c := NewCommentAction(CommentActionParam{CommentCommand: model})
		action, err := c.CommentAction(ctx, delReq)
		assertions.NoError(err)
		assertions.NotNil(action)
	})
	t.Run("delete comment failed", func(t *testing.T) {
		model := NewCommentCommand(t)
		c := NewCommentAction(CommentActionParam{CommentCommand: model})
		model.On("Delete", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed"))
		action, err := c.CommentAction(ctx, delReq)
		assertions.Error(err)
		assertions.Nil(action)
	})
	t.Run("invalid action type", func(t *testing.T) {
		c := NewCommentAction(CommentActionParam{})
		action, err := c.CommentAction(ctx, &types.CommentActionReq{
			ActionType: 3,
		})
		assertions.Error(err)
		assertions.Nil(action)
	})
}

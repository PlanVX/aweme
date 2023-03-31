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
	"testing"
)

func TestFollow(t *testing.T) {

	// create a followAction logic with the mock followAction model
	followAction := NewFollowAction(FollowActionParam{RelationCommand: nil})

	t.Run("follow success", func(t *testing.T) {
		// create a mock followAction model
		relationCommand := NewRelationCommand(t)
		followAction.relationCommand = relationCommand
		// create a context with a mock value
		ctx := ContextWithOwner(int64(456))

		// create a followAction action request to follow user with ID 123
		req := &types.RelationActionReq{
			ActionType: 1,
			ToUserID:   123,
		}

		// define the expected output
		expected := &types.RelationActionResp{}

		// mock the Insert function of the followAction model to return nil
		relationCommand.On("Insert", ctx, &dal.Relation{UserID: int64(456), FollowTo: int64(123)}).Return(nil)

		// call the Follow function and check the output and error
		resp, err := followAction.Follow(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, expected, resp)

		// assert that the Insert function of the followAction model was called once with the correct arguments
		relationCommand.AssertCalled(t, "Insert", ctx, &dal.Relation{UserID: int64(456), FollowTo: int64(123)})

	})
	t.Run("follow error", func(t *testing.T) {
		// create a mock followAction model
		relationCommand := NewRelationCommand(t)
		followAction.relationCommand = relationCommand
		// create a context with a mock value
		ctx := ContextWithOwner(int64(456))

		// create a followAction action request to follow user with ID 123
		req := &types.RelationActionReq{
			ActionType: 1,
			ToUserID:   123,
		}

		// mock the Insert function of the followAction model to return an error
		relationCommand.On("Insert", ctx, &dal.Relation{UserID: int64(456), FollowTo: int64(123)}).Return(errors.New("error"))

		// call the Follow function and check the output and error
		resp, err := followAction.Follow(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)

		// assert that the Insert function of the followAction model was called once with the correct arguments
		relationCommand.AssertCalled(t, "Insert", ctx, &dal.Relation{UserID: int64(456), FollowTo: int64(123)})
	})

}

func TestUnfollow(t *testing.T) {

	// create a followAction logic with the mock followAction model
	followAction := NewFollowAction(FollowActionParam{RelationCommand: nil})

	t.Run("unfollow success", func(t *testing.T) {
		// create a mock followAction model
		relationCommand := NewRelationCommand(t)
		followAction.relationCommand = relationCommand
		// create a context with a mock value
		ctx := ContextWithOwner(int64(456))

		// create a followAction action request to unfollow user with ID 123
		req := &types.RelationActionReq{
			ActionType: 2,
			ToUserID:   123,
		}

		// define the expected output
		expected := &types.RelationActionResp{}

		// mock the Delete function of the followAction model to return nil
		relationCommand.On("Delete", ctx, int64(456), int64(123)).Return(nil)

		// call the Unfollow function and check the output and error
		resp, err := followAction.Follow(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, expected, resp)

		// assert that the Delete function of the followAction model was called once with the correct arguments
		relationCommand.AssertCalled(t, "Delete", ctx, int64(456), int64(123))
	})
	t.Run("unfollow error", func(t *testing.T) {
		// create a mock followAction model
		relationModel := NewRelationCommand(t)
		followAction.relationCommand = relationModel
		// create a context with a mock value
		ctx := ContextWithOwner(int64(456))

		// create a followAction action request to unfollow user with ID 123
		req := &types.RelationActionReq{
			ActionType: 2,
			ToUserID:   123,
		}

		// mock the Delete function of the followAction model to return an error
		relationModel.On("Delete", ctx, int64(456), int64(123)).Return(errors.New("error"))

		// call the Unfollow function and check the output and error
		resp, err := followAction.Follow(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)

		// assert that the Delete function of the followAction model was called once with the correct arguments
		relationModel.AssertCalled(t, "Delete", ctx, int64(456), int64(123))
	})
}

func TestRelationAction_Follow(t *testing.T) {
	// create a relation logic with the mock relation model
	relation := NewFollowAction(FollowActionParam{RelationCommand: nil})

	// create a context with a mock value
	ctx := ContextWithOwner(int64(456))

	// create a relation action request to follow user with unknown action type
	req := &types.RelationActionReq{
		ActionType: 3,
		ToUserID:   123,
	}
	// call the Follow function and check the output and error
	_, err := relation.Follow(ctx, req)
	assert.Error(t, err)
}

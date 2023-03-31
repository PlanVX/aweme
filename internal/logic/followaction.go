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
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type (
	// FollowAction is the relation logic layer struct
	FollowAction struct {
		relationCommand dal.RelationCommand
	}
	// FollowActionParam is the param for NewFollowAction
	FollowActionParam struct {
		fx.In
		RelationCommand dal.RelationCommand
	}
)

// NewFollowAction returns a new FollowAction logic
func NewFollowAction(param FollowActionParam) *FollowAction {
	return &FollowAction{
		relationCommand: param.RelationCommand,
	}
}

// Follow handles the follow action
func (r *FollowAction) Follow(c context.Context, req *types.RelationActionReq) (*types.RelationActionResp, error) {
	follower, _ := c.Value(ContextKey).(int64) // get the follower from context

	switch req.ActionType {
	case int32(1): // means follow a user
		rel := &dal.Relation{
			UserID:   follower,
			FollowTo: req.ToUserID,
		}
		err := r.relationCommand.Insert(c, rel)
		if err != nil {
			return nil, err
		}
	case int32(2): // means unfollow a user
		err := r.relationCommand.Delete(c, follower, req.ToUserID)
		if err != nil {
			return nil, err
		}
	default:
		return nil, echo.ErrBadRequest
	}

	return &types.RelationActionResp{}, nil
}

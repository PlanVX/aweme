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
	// Like is the like logic layer struct
	Like struct {
		likeCommand dal.LikeCommand
	}
	// LikeParam is the param for NewLike
	LikeParam struct {
		fx.In
		LikeCommand dal.LikeCommand
	}
)

// NewLike returns a new Like logic
func NewLike(param LikeParam) *Like {
	return &Like{
		likeCommand: param.LikeCommand,
	}
}

// Like is the like logic
// handle the like action
func (l *Like) Like(c context.Context, req *types.FavoriteActionReq) (*types.FavoriteActionResp, error) {
	owner, _ := c.Value(ContextKey).(int64) // get the owner from context
	switch req.ActionType {
	case int32(1): // means add like for a video

		like := &dal.Like{
			VideoID: req.VideoID,
			UserID:  owner,
		}

		err := l.likeCommand.Insert(c, like)
		if err != nil {
			return nil, err
		}
	case int32(2): // means remove like for a video

		err := l.likeCommand.Delete(c, req.VideoID, owner)
		if err != nil {
			return nil, err
		}
	default:
		return nil, echo.ErrBadRequest
	}
	return &types.FavoriteActionResp{}, nil
}

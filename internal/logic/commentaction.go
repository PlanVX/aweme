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
	// CommentAction is the comment logic layer struct
	CommentAction struct {
		userQuery      dal.UserQuery
		commentCommand dal.CommentCommand
	}
	// CommentActionParam is the parameter for NewCommentAction
	CommentActionParam struct {
		fx.In
		UserQuery      dal.UserQuery
		CommentCommand dal.CommentCommand
	}
)

// NewCommentAction returns a new CommentAction logic
func NewCommentAction(param CommentActionParam) *CommentAction {
	return &CommentAction{userQuery: param.UserQuery, commentCommand: param.CommentCommand}
}

// CommentAction 评论逻辑
func (c *CommentAction) CommentAction(ctx context.Context, req *types.CommentActionReq) (*types.CommentActionResp, error) {
	// 首先获取userid（登录用户id）
	userid, _ := ctx.Value(ContextKey).(int64)

	if req.ActionType == 1 { //新增评论

		comment := &dal.Comment{ // 创建评论
			Content: req.CommentText,
			VideoID: req.VideoID,
			UserID:  userid,
		}

		err := c.commentCommand.Insert(ctx, comment)
		if err != nil {
			return nil, err
		}

		user, err := c.userQuery.FindOne(ctx, userid)
		if err != nil {
			return nil, err
		}

		// 评论成功注意返回评论内容
		return &types.CommentActionResp{
			Comment: &types.Comment{
				ID:         comment.ID,
				Content:    comment.Content,
				CreateDate: comment.CreatedAt.Format("01-02"),
				User:       covertUser(user),
			},
		}, nil

	} else if req.ActionType == 2 {

		err := c.commentCommand.Delete(ctx, req.CommentID, userid, req.VideoID)
		if err != nil {
			return nil, err
		}

		return &types.CommentActionResp{}, nil

	} else {
		return nil, echo.ErrBadRequest
	}
}

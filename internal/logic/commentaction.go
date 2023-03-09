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
		userModel    dal.UserModel
		commentModel dal.CommentModel
	}
	// CommentActionParam is the parameter for NewCommentAction
	CommentActionParam struct {
		fx.In
		UserModel    dal.UserModel
		CommentModel dal.CommentModel
	}
)

// NewCommentAction returns a new CommentAction logic
func NewCommentAction(param CommentActionParam) *CommentAction {
	return &CommentAction{userModel: param.UserModel, commentModel: param.CommentModel}
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

		err := c.commentModel.Insert(ctx, comment)
		if err != nil {
			return nil, err
		}

		user, err := c.userModel.FindOne(ctx, userid)
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

		err := c.commentModel.Delete(ctx, req.CommentID, userid, req.VideoID)
		if err != nil {
			return nil, err
		}

		return &types.CommentActionResp{}, nil

	} else {
		return nil, echo.ErrBadRequest
	}
}

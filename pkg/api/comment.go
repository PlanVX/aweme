package api

import (
	"context"
	"github.com/PlanVX/aweme/pkg/types"
)

// NewCommentAction godoc
// @Summary 评论操作
// @Description 评论操作
// @Tags 互动接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user formData types.CommentActionReq true "评论信息"
// @Success 200 {object} types.CommentActionResp
// @Router /comment/action/ [post]
func NewCommentAction() *Api {
	return &Api{
		Method: "POST",
		Path:   "/comment/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.CommentActionReq) (*types.CommentActionResp, error) {
			return nil, nil
		}),
	}
}

// NewCommentList godoc
// @Summary 评论列表
// @Description 评论列表
// @Tags 互动接口
// @Produce json
// @Param user_id query types.CommentListReq true "用户信息"
// @Success 200 {object} types.CommentListResp
// @Router /comment/list/ [get]
func NewCommentList() *Api {
	return &Api{
		Method: "GET",
		Path:   "/comment/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.CommentListReq) (*types.CommentListResp, error) {
			return nil, nil
		}),
	}
}

package api

import (
	"github.com/PlanVX/aweme/pkg/logic"
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
func NewCommentAction(action *logic.CommentAction) *Api {
	return &Api{
		Method:  "POST",
		Path:    "/comment/action/",
		Handler: WrapperFunc(action.CommentAction),
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
func NewCommentList(list *logic.CommentList) *Api {
	return &Api{
		Method:  "GET",
		Path:    "/comment/list/",
		Handler: WrapperFunc(list.CommentList),
	}
}

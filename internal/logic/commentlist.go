package logic

import (
	"context"

	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"go.uber.org/fx"
)

type (
	// CommentList is the comment logic layer struct
	CommentList struct {
		userQuery     dal.UserQuery
		commentQuery  dal.CommentQuery
		relationQuery dal.RelationQuery
	}
	// CommentListParam is the parameter for NewCommentList
	CommentListParam struct {
		fx.In
		UserQuery     dal.UserQuery
		CommentQuery  dal.CommentQuery
		RelationQuery dal.RelationQuery
	}
)

// NewCommentList returns a new CommentList logic
func NewCommentList(param CommentListParam) *CommentList {
	return &CommentList{userQuery: param.UserQuery, commentQuery: param.CommentQuery, relationQuery: param.RelationQuery}
}

// CommentList 评论列表逻辑
func (c *CommentList) CommentList(ctx context.Context, req *types.CommentListReq) (resp *types.CommentListResp, err error) {

	// 首先获取userid（登录用户id）
	userid, _ := ctx.Value(ContextKey).(int64)

	commentList, err := c.commentQuery.FindByVideoID(ctx, req.VideoID, 30, 0)
	if err != nil {
		return nil, err
	}

	// 获取评论对应的用户id列表
	userIds := extractUserIDsFromComment(commentList)

	// 获取用户列表
	userList, err := c.userQuery.FindMany(ctx, userIds)
	if err != nil {
		return nil, err
	}

	var list []int64
	if userid != 0 {
		// 获取关注关系列表
		list, err = c.relationQuery.FindWhetherFollowedList(ctx, userid, userIds)
		if err != nil {
			return nil, err
		}
	}

	// 转换为CommentListResp
	commentListObjs := packComments(commentList, userList, list)

	// 找到评论列表之后
	return &types.CommentListResp{
		StatusCode:  0,
		StatusMsg:   "获取评论列表成功",
		CommentList: commentListObjs,
	}, nil
}

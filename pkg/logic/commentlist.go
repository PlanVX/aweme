package logic

import (
	"context"

	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type (
	// CommentList is the comment logic layer struct
	CommentList struct {
		userModel     dal.UserModel
		commentModel  dal.CommentModel
		relationModel dal.RelationModel
	}
	// CommentListParam is the parameter for NewCommentList
	CommentListParam struct {
		fx.In
		UserModel     dal.UserModel
		CommentModel  dal.CommentModel
		RelationModel dal.RelationModel
	}
)

// NewCommentList returns a new CommentList logic
func NewCommentList(param CommentListParam) *CommentList {
	return &CommentList{userModel: param.UserModel, commentModel: param.CommentModel, relationModel: param.RelationModel}
}

// CommentList 评论列表逻辑
func (c *CommentList) CommentList(ctx context.Context, req *types.CommentListReq) (resp *types.CommentListResp, err error) {

	// 首先获取userid（登录用户id）
	userid, _ := ctx.Value(ContextKey).(int64)

	commentList, err := c.commentModel.FindByVideoID(ctx, req.VideoID, 30, 0)
	if err != nil {
		return nil, err
	}
	// 获取评论对应的用户id列表
	userIds := lo.Map(commentList, func(comment *dal.Comment, _ int) int64 {
		return comment.UserID
	})
	// 获取用户列表
	userList, err := c.userModel.FindMany(ctx, userIds)
	if err != nil {
		return nil, err
	}
	// 转换为map
	userMappings := lo.SliceToMap(userList, func(user *dal.User) (int64, *types.User) {
		return user.ID, covertUser(user)
	})
	// 获取关注关系列表
	list, err := c.relationModel.FindWhetherFollowedList(ctx, userid, userIds)
	if err != nil {
		return nil, err
	}
	// 转换为map
	followMappings := lo.SliceToMap(list, func(relation int64) (int64, bool) { return relation, true })
	// 转换为CommentListResp
	commentListObjs := lo.Map(commentList, func(comment *dal.Comment, _ int) *types.Comment {
		user, _ := userMappings[comment.UserID]
		user.IsFollow, _ = followMappings[comment.UserID] // 不存在则默认为false
		return &types.Comment{
			ID:         comment.ID,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
			User:       user,
		}
	})
	// 找到评论列表之后
	return &types.CommentListResp{
		StatusCode:  0,
		StatusMsg:   "获取评论列表成功",
		CommentList: commentListObjs,
	}, nil
}

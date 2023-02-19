package logic

import (
	"context"

	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type (
	// Login is the logic for login
	CommentList struct {
		userModel     dal.UserModel
		commentModel  dal.CommentModel
		relationModel dal.RelationModel
	}
	// LoginParam is the parameter for NewLogin
	CommentListParam struct {
		fx.In
		userModel     dal.UserModel
		commentModel  dal.CommentModel
		relationModel dal.RelationModel
	}
)

// NewPublishList returns a new PublishList logic
func NewCommentList(param CommentListParam) *CommentList {
	return &CommentList{userModel: param.userModel, commentModel: param.commentModel, relationModel: param.relationModel}
}

// 查看视频评论列表逻辑
func (c *CommentList) CommentList(ctx context.Context, req *types.CommentListReq) (resp *types.CommentListResp, err error) {

	// 首先获取userid（登录用户id）
	userid, _ := ctx.Value(ContextKey).(int64)

	commentList, err := c.commentModel.FindByVideoID(ctx, req.VideoID, 30, 0)
	if err != nil {
		return &types.CommentListResp{
			StatusCode: -1,
			StatusMsg:  "获取评论列表失败",
		}, err
	}
	commentListObjs := lo.Map(commentList, func(comment *dal.Comment, _ int) *types.Comment {
		// 获取用户信息
		user, err := c.userModel.FindOne(ctx, comment.UserID)
		if err != nil {
			// 查找用户失败
			user = &dal.User{}
		}
		followcount, err := c.relationModel.GetNumByFollowerID(ctx, user.ID)
		if err != nil {
			// 查询关注人数失败
			followcount = 0
		}
		followercount, err := c.relationModel.GetNumByFollowedID(ctx, user.ID)
		if err != nil {
			// 查询粉丝人数失败
			followercount = 0
		}
		followrelcount, err := c.relationModel.GetNumByFollowerIDAndFollowedID(ctx, userid, user.ID)
		if err != nil {
			followrelcount = 0
		}
		isfollow := false
		if followrelcount > 0 {
			isfollow = true
		}
		return &types.Comment{
			ID:         comment.ID,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
			User: &types.User{
				ID:            comment.UserID,
				Username:      user.Username,
				Avatar:        user.Avatar,
				FollowCount:   int64(followcount),
				FollowerCount: int64(followercount),
				IsFollow:      isfollow,
			},
		}
	})
	// 找到评论列表之后
	return &types.CommentListResp{
		StatusCode:  0,
		StatusMsg:   "获取评论列表成功",
		CommentList: commentListObjs,
	}, nil
}

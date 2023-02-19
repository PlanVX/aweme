package logic

import (
	"context"
	"time"

	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"go.uber.org/fx"
)

type (
	// Login is the logic for login
	CommentAction struct {
		relationModel dal.RelationModel
		userModel     dal.UserModel
		videoModel    dal.VideoModel
		commentModel  dal.CommentModel
	}
	// LoginParam is the parameter for NewLogin
	CommentActionParam struct {
		fx.In
		relationModel dal.RelationModel
		userModel     dal.UserModel
		videoModel    dal.VideoModel
		commentModel  dal.CommentModel
	}
)

// NewPublishList returns a new PublishList logic
func NewCommentAction(param CommentActionParam) *CommentAction {
	return &CommentAction{relationModel: param.relationModel, userModel: param.userModel, videoModel: param.videoModel, commentModel: param.commentModel}
}

// 评论逻辑
func (c *CommentAction) CommentAction(ctx context.Context, req *types.CommentActionReq) (resp *types.CommentActionResp, err error) {

	// 首先获取userid（登录用户id）
	userid, _ := ctx.Value(ContextKey).(int64)

	if err != nil {
		// 报文+登陆错误消息
		return &types.CommentActionResp{
			StatusCode: -2,
			StatusMsg:  "登陆失败",
		}, err
	}
	if req.ActionType == 1 {
		//新增评论
		//创建评论对象
		comment := &dal.Comment{
			Content:   req.CommentText,
			VideoID:   req.VideoID,
			UserID:    userid,
			CreatedAt: time.Now(),
		}
		err = c.commentModel.Insert(ctx, comment)
		if err != nil {
			// 评论失败信息
			return &types.CommentActionResp{
				StatusCode: -1,
				StatusMsg:  "评论失败",
			}, err
		}
		user, err := c.userModel.FindOne(ctx, userid)
		if err != nil {
			// 查找用户失败
			return &types.CommentActionResp{
				StatusCode: -1,
				StatusMsg:  "查找用户失败",
			}, err
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
		// 评论成功注意返回评论内容
		return &types.CommentActionResp{
			StatusCode: 0,
			StatusMsg:  "评论成功",
			Comment: &types.Comment{
				ID:         comment.ID,
				Content:    comment.Content,
				CreateDate: string(comment.CreatedAt.Format("01-02")),
				User: &types.User{
					ID:            comment.UserID,
					Username:      user.Username,
					Avatar:        user.Avatar,
					FollowCount:   int64(followcount),
					FollowerCount: int64(followercount),
					IsFollow:      false,
				},
			},
		}, nil
	}
	if req.ActionType == 2 {
		// 删除评论
		err = c.commentModel.Delete(ctx, req.CommentID, userid)
		if err != nil {
			// 删除评论失败信息
			return &types.CommentActionResp{
				StatusCode: -1,
				StatusMsg:  "删除评论失败",
			}, err
		}
		return &types.CommentActionResp{
			StatusCode: 0,
			StatusMsg:  "删除评论成功",
		}, nil
	}
	return nil, nil
}

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
	FavoriteAction struct {
		videoModel   dal.VideoModel
		commentModel dal.CommentModel
		likeModel    dal.LikeModel
		signer       *JWTSigner
	}
	// LoginParam is the parameter for NewLogin
	FavoriteActionParam struct {
		fx.In
		videoModel   dal.VideoModel
		commentModel dal.CommentModel
		likeModel    dal.LikeModel
		signer       *JWTSigner
	}
)

// NewPublishList returns a new PublishList logic
func NewFavoriteAction(param FavoriteActionParam) *FavoriteAction {
	return &FavoriteAction{likeModel: param.likeModel, videoModel: param.videoModel, commentModel: param.commentModel, signer: param.signer}
}

func (f *FavoriteAction) FavoriteAction(ctx context.Context, req *types.FavoriteActionReq) (resp *types.FavoriteActionResp, err error) {
	//给视频点赞
	// 首先获取userid（登录用户id）
	userid, err := f.signer.parseUserID(req.Token)
	if err != nil {
		// 报文+登陆错误消息
		return &types.FavoriteActionResp{
			StatusCode: -2,
			StatusMsg:  "登陆失败",
		}, err
	}
	// 点赞逻辑
	if req.ActionType == 1 {
		like := &dal.Like{
			VideoID:   req.VideoID,
			UserID:    userid,
			CreatedAt: time.Now(),
		}
		err = f.likeModel.Insert(ctx, like)
		if err != nil {
			// 点赞失败
			return &types.FavoriteActionResp{
				StatusCode: -1,
				StatusMsg:  "点赞失败",
			}, err
		}
		return &types.FavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		}, nil
	}
	// 取消点赞逻辑
	if req.ActionType == 2 {
		err = f.likeModel.Delete(ctx, req.VideoID, userid)
		if err != nil {
			// 点赞失败
			return &types.FavoriteActionResp{
				StatusCode: -1,
				StatusMsg:  "取消点赞失败",
			}, err
		}
		return &types.FavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "取消点赞成功",
		}, nil
	}
	// 错误逻辑
	return nil, nil
}

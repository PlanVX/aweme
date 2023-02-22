package logic

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"go.uber.org/fx"
)

type (
	// LikeList is the like list logic layer struct
	LikeList struct {
		likeModel  dal.LikeModel
		videoModel dal.VideoModel
		userModel  dal.UserModel
	}
	// LikeListParam is the param for NewLikeList
	LikeListParam struct {
		fx.In
		LikeModel  dal.LikeModel
		UserModel  dal.UserModel
		VideoModel dal.VideoModel
	}
)

// NewLikeList returns a new LikeList logic
func NewLikeList(param LikeListParam) *LikeList {
	return &LikeList{
		likeModel:  param.LikeModel,
		userModel:  param.UserModel,
		videoModel: param.VideoModel,
	}
}

// LikeList is the like list logic
// handle the like list
func (l *LikeList) LikeList(ctx context.Context, req *types.FavoriteListReq) (*types.FavoriteListResp, error) {
	owner, _ := ctx.Value(ContextKey).(int64)
	likes, err := l.likeModel.FindVideoIDsByUserID(ctx, req.UserID, 30, 0)
	if err != nil {
		return nil, err
	}
	many, err := l.videoModel.FindMany(ctx, likes)
	if err != nil {
		return nil, err
	}
	users, err := l.userModel.FindMany(ctx, extractUserIDs(many)) // 根据用户 id 批量查询用户信息
	if err != nil {
		return nil, err
	}
	likedList, err := l.likeModel.FindWhetherLiked(ctx, owner, likes)
	if err != nil {
		return nil, err
	}
	// set is favorite
	videos := packVideos(many, users, likedList)
	return &types.FavoriteListResp{
		VideoList: videos,
	}, nil
}

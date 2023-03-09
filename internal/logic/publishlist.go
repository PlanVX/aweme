package logic

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"go.uber.org/fx"
)

type (
	// PublishList is the logic for publish list
	PublishList struct {
		videoModel dal.VideoModel
		userModel  dal.UserModel
		likeModel  dal.LikeModel
	}

	// PublishListParam is the parameter for NewPublishList
	PublishListParam struct {
		fx.In
		VideoModel dal.VideoModel
		UserModel  dal.UserModel
		LikeModel  dal.LikeModel
	}
)

// NewPublishList returns a new PublishList logic
func NewPublishList(param PublishListParam) *PublishList {
	return &PublishList{
		videoModel: param.VideoModel,
		userModel:  param.UserModel,
		likeModel:  param.LikeModel,
	}
}

// PublishList gets publish list
func (p *PublishList) PublishList(ctx context.Context, req *types.PublishListReq) (*types.PublishListResp, error) {

	// query videos by specified user id
	filteredVideos, err := p.videoModel.FindByUserID(ctx, req.UserID, 30, 0)
	if err != nil {
		return nil, err
	}

	// query user info by specified user id
	user, err := p.userModel.FindOne(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	// video id list
	videoIDs := extractVideosIDs(filteredVideos)

	// query like info by specified user id
	likes, err := p.likeModel.FindWhetherLiked(ctx, req.UserID, videoIDs)
	if err != nil {
		return nil, err
	}

	videos := packVideos(filteredVideos, []*dal.User{user}, likes)

	return &types.PublishListResp{VideoList: videos}, nil
}

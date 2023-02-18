package logic

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type (
	// PublishList is the logic for publish list
	PublishList struct {
		videoModel dal.VideoModel
		userModel  dal.UserModel
	}

	// PublishListParam is the parameter for NewPublishList
	PublishListParam struct {
		fx.In
		VideoModel dal.VideoModel
		UserModel  dal.UserModel
	}
)

// NewPublishList returns a new PublishList logic
func NewPublishList(param PublishListParam) *PublishList {
	return &PublishList{videoModel: param.VideoModel, userModel: param.UserModel}
}

// PublishList gets publish list
func (p *PublishList) PublishList(ctx context.Context, req *types.PublishListReq) (*types.PublishListResp, error) {
	// query videos by specified user id
	filteredVideos, err := p.videoModel.FindByUserID(ctx, req.UserID, 30, 0)
	if err != nil {
		return nil, err
	}
	user, err := p.userModel.FindOne(ctx, req.UserID) // query user info by specified user id
	if err != nil {
		return nil, err
	}
	videos := lo.Map(filteredVideos, func(video *dal.Video, _ int) *types.Video {
		v := &types.Video{ID: video.ID, Title: video.Title, CoverURL: video.CoverURL, PlayURL: video.VideoURL}
		v.Author = &types.User{ID: user.ID, Username: user.Username, Avatar: user.Avatar}
		return v
	}) // convert dal.Video to types.Video and fill author info
	return &types.PublishListResp{VideoList: videos}, nil
}

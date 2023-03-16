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
		videoQuery dal.VideoQuery
		userQuery  dal.UserQuery
		likeQuery  dal.LikeQuery
	}

	// PublishListParam is the parameter for NewPublishList
	PublishListParam struct {
		fx.In
		VideoQuery dal.VideoQuery
		UserQuery  dal.UserQuery
		LikeQuery  dal.LikeQuery
	}
)

// NewPublishList returns a new PublishList logic
func NewPublishList(param PublishListParam) *PublishList {
	return &PublishList{
		videoQuery: param.VideoQuery,
		userQuery:  param.UserQuery,
		likeQuery:  param.LikeQuery,
	}
}

// PublishList gets publish list
func (p *PublishList) PublishList(ctx context.Context, req *types.PublishListReq) (*types.PublishListResp, error) {

	// query videos by specified user id
	filteredVideos, err := p.videoQuery.FindByUserID(ctx, req.UserID, 30, 0)
	if err != nil {
		return nil, err
	}

	// query user info by specified user id
	user, err := p.userQuery.FindOne(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	// video id list
	videoIDs := extractVideosIDs(filteredVideos)

	// query like info by specified user id
	likes, err := p.likeQuery.FindWhetherLiked(ctx, req.UserID, videoIDs)
	if err != nil {
		return nil, err
	}

	videos := packVideos(filteredVideos, []*dal.User{user}, likes)

	return &types.PublishListResp{VideoList: videos}, nil
}

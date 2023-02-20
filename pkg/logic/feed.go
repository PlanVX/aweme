package logic

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"go.uber.org/fx"
)

type (
	// Feed is the logic for feed
	Feed struct {
		videoModel dal.VideoModel
		userModel  dal.UserModel
		likeModel  dal.LikeModel
	}
	// FeedParam is the parameter for NewFeed
	FeedParam struct {
		fx.In
		VideoModel dal.VideoModel
		UserModel  dal.UserModel
		LikeModel  dal.LikeModel
	}
)

// NewFeed creates a new feed logic
func NewFeed(param FeedParam) *Feed {
	return &Feed{
		videoModel: param.VideoModel,
		userModel:  param.UserModel,
		likeModel:  param.LikeModel,
	}
}

// Feed 获取首页视频流
func (f *Feed) Feed(ctx context.Context, req *types.FeedReq) (resp *types.FeedResp, err error) {
	owner, _ := ctx.Value(ContextKey).(int64)
	latestVideo, err := f.videoModel.FindLatest(ctx, req.LatestTime, 30) // 查询 30 个视频
	if err != nil {
		return nil, err
	}
	users, err := f.userModel.FindMany(ctx, extractUserIDs(latestVideo)) // 根据用户 id 批量查询用户信息
	if err != nil {
		return nil, err
	}
	videoIDs := extractVideosIDs(latestVideo)
	likedList, err := f.likeModel.FindWhetherLiked(ctx, owner, videoIDs)
	if err != nil {
		return nil, err
	}
	videos := packVideos(latestVideo, users, likedList)
	return &types.FeedResp{
		VideoList: videos,
	}, nil
}

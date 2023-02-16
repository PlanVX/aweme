package logic

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type (
	// Feed is the logic for feed
	Feed struct {
		videoModel dal.VideoModel
		userModel  dal.UserModel
	}
	// FeedParam is the parameter for NewFeed
	FeedParam struct {
		fx.In
		VideoModel dal.VideoModel
		UserModel  dal.UserModel
	}
)

// NewFeed creates a new feed logic
func NewFeed(param FeedParam) *Feed {
	return &Feed{
		videoModel: param.VideoModel,
		userModel:  param.UserModel,
	}
}

// Feed 获取首页视频流
func (f *Feed) Feed(ctx context.Context, req *types.FeedReq) (resp *types.FeedResp, err error) {
	latestVideo, err := f.videoModel.FindLatest(ctx, req.LatestTime, 30) // 查询 30 个视频
	if err != nil {
		return nil, err
	}
	videos := lo.Map(latestVideo, func(v *dal.Video, _ int) *types.Video {
		return &types.Video{CoverURL: v.CoverURL, ID: v.ID, PlayURL: v.VideoURL, Title: v.Title}
	}) // 将视频列表转换为 types.Video 列表
	uids := lo.Map(latestVideo, func(v *dal.Video, _ int) int64 { return v.UserID }) // 获取作者用户 id 列表
	users, err := f.userModel.FindMany(ctx, uids)                                    // 根据用户 id 批量查询用户信息
	if err != nil {
		return nil, err
	}
	mappings := lo.SliceToMap(users, func(u *dal.User) (int64, *types.User) {
		return u.ID, &types.User{Avatar: u.Avatar, ID: u.ID, Username: u.Username}
	}) // 将用户列表转换为 map
	lo.ForEach(videos, func(v *types.Video, _ int) {
		v.Author, _ = mappings[v.ID]
	}) // 为每个视频填充用户信息
	return &types.FeedResp{
		VideoList: videos,
	}, nil
}

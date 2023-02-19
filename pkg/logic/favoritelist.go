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
	FavoriteList struct {
		userModel     dal.UserModel
		videoModel    dal.VideoModel
		commentModel  dal.CommentModel
		likeModel     dal.LikeModel
		relationModel dal.RelationModel
	}
	// LoginParam is the parameter for NewLogin
	FavoriteListParam struct {
		fx.In
		userModel     dal.UserModel
		videoModel    dal.VideoModel
		commentModel  dal.CommentModel
		likeModel     dal.LikeModel
		relationModel dal.RelationModel
	}
)

// NewPublishList returns a new PublishList logic
func NewFavoriteList(param FavoriteListParam) *FavoriteList {
	return &FavoriteList{userModel: param.userModel, likeModel: param.likeModel, videoModel: param.videoModel, commentModel: param.commentModel, relationModel: param.relationModel}
}

// 登录用户的所有点赞视频逻辑
func (f *FavoriteList) FavoriteList(ctx context.Context, req *types.FavoriteListReq) (resp *types.FavoriteListResp, err error) {
	likelist, err := f.likeModel.FindByUserID(ctx, req.UserID)
	if err != nil {
		return &types.FavoriteListResp{
			StatusCode: -1,
			StatusMsg:  "获取点赞视频id失败",
		}, err
	}
	ids := lo.Map(likelist, func(like *dal.Like, _ int) int64 {
		return like.VideoID
	})
	videolist, err := f.videoModel.FindMany(ctx, ids)
	if err != nil {
		return &types.FavoriteListResp{
			StatusCode: -1,
			StatusMsg:  "获取点赞视频失败",
		}, err
	}
	// 要根据video信息查出types.Video信息
	// commentcount 评论总数
	videoListObjs := lo.Map(videolist, func(video *dal.Video, _ int) *types.Video {
		commentnum, err := f.commentModel.GetNumByVideoID(ctx, video.ID)
		if err != nil {
			//查询评论总数出错，默认显示为0
			commentnum = 0
		}
		likenum, err := f.likeModel.GetNumByVideoID(ctx, video.ID)
		if err != nil {
			likenum = 0
		}
		// 查询用户是否点赞
		likecount, err := f.likeModel.GetNumByVideoIDAndUserID(ctx, video.ID, req.UserID)
		if err != nil {
			likecount = 0
		}
		liked := false
		if likecount > 0 {
			liked = true
		}
		user, err := f.userModel.FindOne(ctx, video.UserID)
		if err != nil {
			// 查询作者失败
			user = &dal.User{}
		}
		followcount, err := f.relationModel.GetNumByFollowerID(ctx, user.ID)
		if err != nil {
			// 查询关注人数失败
			followcount = 0
		}
		followercount, err := f.relationModel.GetNumByFollowedID(ctx, user.ID)
		if err != nil {
			// 查询粉丝人数失败
			followercount = 0
		}
		followrelcount, err := f.relationModel.GetNumByFollowerIDAndFollowedID(ctx, req.UserID, user.ID)
		if err != nil {
			followrelcount = 0
		}
		isfollow := false
		if followrelcount > 0 {
			isfollow = true
		}
		userobj := &types.User{
			ID:            user.ID,
			Username:      user.Username,
			Avatar:        user.Avatar,
			FollowCount:   int64(followcount),
			FollowerCount: int64(followercount),
			IsFollow:      isfollow,
		}
		return &types.Video{
			ID:            video.ID,
			Author:        userobj,
			CommentCount:  int64(commentnum),
			CoverURL:      video.CoverURL,
			FavoriteCount: int64(likenum),
			IsFavorite:    liked,
		}
	})
	return &types.FavoriteListResp{
		StatusCode: 0,
		StatusMsg:  "获取点赞视频成功",
		VideoList:  videoListObjs,
	}, nil
}

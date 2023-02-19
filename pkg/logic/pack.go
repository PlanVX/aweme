package logic

import (
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"github.com/samber/lo"
)

// composite dal.Video list and dal.User list to types.Video list
func packVideos(videos []*dal.Video, users []*dal.User) []*types.Video {
	base := lo.Map(videos, func(v *dal.Video, _ int) *types.Video {
		return &types.Video{CoverURL: v.CoverURL, ID: v.ID, PlayURL: v.VideoURL, Title: v.Title}
	}) // 将视频列表转换为 types.Video 列表
	mappings := lo.SliceToMap(users, func(u *dal.User) (int64, *types.User) {
		return u.ID, &types.User{Avatar: u.Avatar, ID: u.ID, Username: u.Username}
	}) // 将用户列表转换为 map
	lo.ForEach(base, func(v *types.Video, i int) {
		uid := videos[i].UserID
		v.Author, _ = mappings[uid]
	})
	return base
}

// extract UserID get user ids of videos
func extractUserIDs(u []*dal.Video) []int64 {
	return lo.Map(u, func(v *dal.Video, _ int) int64 { return v.UserID })
}

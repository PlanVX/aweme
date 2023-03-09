package logic

import (
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/samber/lo"
)

// composite dal.Video list and dal.User list to types.Video list
// liked is the list of liked video id of current user
func packVideos(videos []*dal.Video, users []*dal.User, liked []int64) []*types.Video {
	base := lo.Map(videos, func(v *dal.Video, _ int) *types.Video {
		return covertVideo(v)
	}) // 将视频列表转换为 types.Video 列表

	mappings := lo.SliceToMap(users, func(u *dal.User) (int64, *types.User) {
		return u.ID, covertUser(u)
	}) // 将用户列表转换为 map

	likedMap := idsMap(liked)

	lo.ForEach(base, func(v *types.Video, i int) {
		uid := videos[i].UserID
		v.Author, _ = mappings[uid]
		v.IsFavorite, _ = likedMap[v.ID]
	})

	return base
}

// extractVideosIDs extract video ids from video list
func extractVideosIDs(latestVideo []*dal.Video) []int64 {
	return lo.Map(latestVideo, func(item *dal.Video, index int) int64 {
		return item.ID
	})
}

// extract UserID get user ids of videos
func extractUserIDs(u []*dal.Video) []int64 {
	return lo.Map(u, func(v *dal.Video, _ int) int64 { return v.UserID })
}

// covertUser convert dal.User to types.User
func covertUser(v *dal.User) *types.User {
	return &types.User{
		ID:              v.ID,
		Username:        v.Username,
		Avatar:          v.Avatar,
		BackgroundImage: v.BackgroundImage,
		Signature:       v.Signature,
		TotalFavorites:  v.BeLikedCount,
		WorkCount:       v.VideoCount,
		FavoriteCount:   v.LikeCount,
		FollowCount:     v.FollowCount,
		FollowerCount:   v.FansCount,
	}
}

// covertVideo convert dal.Video to types.Video
func covertVideo(v *dal.Video) *types.Video {
	return &types.Video{
		ID:            v.ID,
		Title:         v.Title,
		CoverURL:      v.CoverURL,
		PlayURL:       v.VideoURL,
		CommentCount:  v.CommentCount,
		FavoriteCount: v.LikeCount,
	}
}

// idsMap is a helper function to convert a id list to a map[int64]bool
func idsMap(followedList []int64) map[int64]bool {
	return lo.SliceToMap(followedList, func(item int64) (int64, bool) { return item, true })
}

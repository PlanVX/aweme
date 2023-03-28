/*
 * Copyright (c) 2023 The PlanVX Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logic

import (
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
)

// composite dal.Video list and dal.User list to types.Video list
// liked is the list of liked video id of current user
func packVideos(videos []*dal.Video, users []*dal.User, liked []int64) (agg []*types.Video) {
	if len(videos) == 0 {
		return agg
	}

	// 将视频列表转换为 types.Video 列表
	agg = make([]*types.Video, len(videos))
	for i, v := range videos {
		agg[i] = covertVideo(v)
	}

	// 将用户列表转换为 map
	mappings := userSliceToMap(users)

	likedMap := idsMap(liked)

	for i, v := range agg {
		uid := videos[i].UserID
		v.Author, _ = mappings[uid]
		v.IsFavorite, _ = likedMap[v.ID]
	}

	return agg
}

func packComments(commentList []*dal.Comment, users []*dal.User, list []int64) (agg []*types.Comment) {

	// 转换为map
	userMappings := userSliceToMap(users)

	// 转换为map
	followMappings := idsMap(list)

	agg = make([]*types.Comment, len(commentList))
	for i, v := range commentList {
		agg[i] = &types.Comment{
			ID:         v.ID,
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
			User:       userMappings[v.UserID],
		}
		agg[i].User.IsFollow, _ = followMappings[v.UserID]
	}
	return agg
}

func userSliceToMap(users []*dal.User) map[int64]*types.User {
	mappings := make(map[int64]*types.User, len(users))
	for _, u := range users {
		mappings[u.ID] = covertUser(u)
	}
	return mappings
}

// extractVideosIDs extract video ids from video list
func extractVideosIDs(videos []*dal.Video) (videoIDs []int64) {
	for _, v := range videos {
		videoIDs = append(videoIDs, v.ID)
	}
	return videoIDs
}

// extract UserID get user ids of videos
func extractUserIDs(videos []*dal.Video) (userIDs []int64) {
	userIDs = make([]int64, len(videos))
	for _, v := range videos {
		userIDs = append(userIDs, v.UserID)
	}
	return userIDs
}

// extractUserIDsFromComment extract user ids from comment list
func extractUserIDsFromComment(comments []*dal.Comment) (userIDs []int64) {
	userIDs = make([]int64, len(comments))
	for _, v := range comments {
		userIDs = append(userIDs, v.UserID)
	}
	return userIDs
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
	m := make(map[int64]bool, len(followedList))
	for _, v := range followedList {
		m[v] = true
	}
	return m
}

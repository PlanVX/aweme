package logic

import (
	"context"
	"errors"
	"testing"

	"github.com/PlanVX/aweme/pkg/types"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/PlanVX/aweme/pkg/dal"
)

func TestLikeList(t *testing.T) {
	assertions := assert.New(t)
	videoIDs := []int64{1, 2}
	dalVideos := []*dal.Video{{ID: 1, UserID: 4}, {ID: 2, UserID: 5}}
	dalUsers := []*dal.User{{ID: 4}, {ID: 5}}
	t.Run("query success", func(t *testing.T) {
		likeModel, videoModel, userModel, list := mockLikeList(t)
		likeModel.On("FindVideoIDsByUserID",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(videoIDs, nil)
		videoModel.On("FindMany", mock.Anything, mock.Anything).Return(dalVideos, nil)
		userModel.On("FindMany", mock.Anything, mock.Anything).Return(dalUsers, nil)

		resp, err := list.LikeList(context.TODO(), &types.FavoriteListReq{
			UserID: 1,
		})
		assertions.NoError(err)
		assertions.NotNil(resp.VideoList)
		lo.ForEach(resp.VideoList, func(item *types.Video, index int) {
			assertions.Equal(dalVideos[index].ID, item.ID)
			assertions.Equal(dalVideos[index].UserID, item.Author.ID)
			assertions.Equal(true, item.IsFavorite)
		})
	})
	t.Run("query like list failed", func(t *testing.T) {
		likeModel, _, _, list := mockLikeList(t)
		likeModel.On("FindVideoIDsByUserID",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(nil, errors.New("query like list failed"))

		_, err := list.LikeList(context.TODO(), &types.FavoriteListReq{
			UserID: 1,
		})
		assertions.Error(err)
	})
	t.Run("query video list failed", func(t *testing.T) {
		likeModel, videoModel, _, list := mockLikeList(t)
		likeModel.On("FindVideoIDsByUserID",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(videoIDs, nil)
		videoModel.On("FindMany", mock.Anything, mock.Anything).Return(nil, errors.New("query video list failed"))

		_, err := list.LikeList(context.TODO(), &types.FavoriteListReq{
			UserID: 1,
		})
		assertions.Error(err)
	})
	t.Run("query user list failed", func(t *testing.T) {
		likeModel, videoModel, userModel, list := mockLikeList(t)
		likeModel.On("FindVideoIDsByUserID",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(videoIDs, nil)
		videoModel.On("FindMany", mock.Anything, mock.Anything).Return(dalVideos, nil)
		userModel.On("FindMany", mock.Anything, mock.Anything).Return(nil, errors.New("query user list failed"))

		_, err := list.LikeList(context.TODO(), &types.FavoriteListReq{
			UserID: 1,
		})
		assertions.Error(err)
	})
}

func mockLikeList(t *testing.T) (*LikeModel, *VideoModel, *UserModel, *LikeList) {
	likeModel := NewLikeModel(t)
	videoModel := NewVideoModel(t)
	userModel := NewUserModel(t)
	list := NewLikeList(LikeListParam{
		LikeModel:  likeModel,
		UserModel:  userModel,
		VideoModel: videoModel,
	})
	return likeModel, videoModel, userModel, list
}

package logic

import (
	"context"
	"errors"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFeed(t *testing.T) {
	assertions := assert.New(t)
	videos := []*dal.Video{{ID: 1, UserID: 1}, {ID: 2, UserID: 2}}
	users := []*dal.User{{ID: 1}, {ID: 2}}
	req := &types.FeedReq{LatestTime: 0}
	liked := []int64{1, 2}
	t.Run("success", func(t *testing.T) {
		u, v, l, feed := mockFeed(t)
		v.On("FindLatest", mock.Anything, mock.Anything, mock.Anything).Return(videos, nil)
		u.On("FindMany", mock.Anything, mock.Anything).Return(users, nil)
		l.On("FindWhetherLiked", mock.Anything, mock.Anything, mock.Anything).Return(liked, nil)
		resp, err := feed.Feed(context.TODO(), req)
		assertions.NoError(err)
		assertions.Equal(2, len(resp.VideoList))
		assertions.Equal(videos[0].ID, resp.VideoList[0].ID)
		assertions.Equal(videos[1].ID, resp.VideoList[1].ID)
		assertions.Equal(users[0].ID, resp.VideoList[0].Author.ID)
		assertions.Equal(users[1].ID, resp.VideoList[1].Author.ID)
		assertions.Equal(true, resp.VideoList[0].IsFavorite)
		assertions.Equal(true, resp.VideoList[1].IsFavorite)
	})
	t.Run("video model error", func(t *testing.T) {
		_, v, _, feed := mockFeed(t)
		v.On("FindLatest", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("video model error"))
		_, err := feed.Feed(context.TODO(), req)
		assertions.Error(err)
	})
	t.Run("user model error", func(t *testing.T) {
		u, v, _, feed := mockFeed(t)
		v.On("FindLatest", mock.Anything, mock.Anything, mock.Anything).Return(videos, nil)
		u.On("FindMany", mock.Anything, mock.Anything).Return(nil, errors.New("user model error"))
		_, err := feed.Feed(context.TODO(), req)
		assertions.Error(err)
	})
	t.Run("like model error", func(t *testing.T) {
		u, v, l, feed := mockFeed(t)
		v.On("FindLatest", mock.Anything, mock.Anything, mock.Anything).Return(videos, nil)
		u.On("FindMany", mock.Anything, mock.Anything).Return(users, nil)
		l.On("FindWhetherLiked", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("like model error"))
		_, err := feed.Feed(context.TODO(), req)
		assertions.Error(err)
	})
}

func mockFeed(t *testing.T) (*UserModel, *VideoModel, *LikeModel, *Feed) {
	u, v, l := NewUserModel(t), NewVideoModel(t), NewLikeModel(t)
	feed := NewFeed(FeedParam{VideoModel: v, UserModel: u, LikeModel: l})
	return u, v, l, feed
}

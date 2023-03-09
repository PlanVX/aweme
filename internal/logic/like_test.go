package logic

import (
	"context"
	"errors"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestLike(t *testing.T) {
	assertions := assert.New(t)
	t.Run("test add like success", func(t *testing.T) {
		model := NewLikeModel(t)
		model.On("Insert", mock.Anything, mock.Anything).Return(nil)
		l := NewLike(LikeParam{LikeModel: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 1, VideoID: 1})
		assertions.NoError(err)
		assertions.NotNil(resp)
	})
	t.Run("test add like failed", func(t *testing.T) {
		model := NewLikeModel(t)
		model.On("Insert", mock.Anything, mock.Anything).Return(errors.New("failed"))
		l := NewLike(LikeParam{LikeModel: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 1, VideoID: 1})
		assertions.Error(err)
		assertions.Nil(resp)
	})
	t.Run("test remove like success", func(t *testing.T) {
		model := NewLikeModel(t)
		model.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		l := NewLike(LikeParam{LikeModel: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 2, VideoID: 1})
		assertions.NoError(err)
		assertions.NotNil(resp)
	})
	t.Run("test remove like failed", func(t *testing.T) {
		model := NewLikeModel(t)
		model.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed"))
		l := NewLike(LikeParam{LikeModel: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 2, VideoID: 1})
		assertions.Error(err)
		assertions.Nil(resp)
	})
	t.Run("test invalid action type", func(t *testing.T) {
		model := NewLikeModel(t)
		l := NewLike(LikeParam{LikeModel: model})
		resp, err := l.Like(context.TODO(), &types.FavoriteActionReq{ActionType: 3, VideoID: 1})
		assertions.Error(err)
		assertions.Nil(resp)
	})
}

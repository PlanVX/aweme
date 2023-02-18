package logic

import (
	"context"
	"errors"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPublishList(t *testing.T) {
	assertions := assert.New(t)
	videos := []*dal.Video{{ID: 1, UserID: 1}, {ID: 2, UserID: 1}, {ID: 3, UserID: 1}}
	user := &dal.User{ID: 1}
	req := &types.PublishListReq{UserID: user.ID}
	t.Run("Success", func(t *testing.T) {
		u, v, list := mockPublishList(t)
		v.On("FindByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(videos, nil)
		u.On("FindOne", mock.Anything, mock.Anything).Return(user, nil)
		resp, err := list.PublishList(context.TODO(), req)
		assertions.NoError(err)
		assertions.Equal(3, len(resp.VideoList))
		assertions.Equal(videos[1].ID, resp.VideoList[1].ID)
		assertions.Equal(resp.VideoList[0].Author.ID, resp.VideoList[1].Author.ID)
		assertions.Equal(user.ID, resp.VideoList[0].Author.ID)
	})
	t.Run("Video model error", func(t *testing.T) {
		_, v, list := mockPublishList(t)
		v.On("FindByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("video model error"))
		_, err := list.PublishList(context.TODO(), req)
		assertions.Error(err)
	})
	t.Run("User model error", func(t *testing.T) {
		u, v, list := mockPublishList(t)
		v.On("FindByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(videos, nil)
		u.On("FindOne", mock.Anything, mock.Anything).Return(nil, errors.New("user model error"))
		_, err := list.PublishList(context.TODO(), req)
		assertions.Error(err)
	})
}

func mockPublishList(t *testing.T) (*UserModel, *VideoModel, *PublishList) {
	u, v := NewUserModel(t), NewVideoModel(t)
	list := NewPublishList(PublishListParam{VideoModel: v, UserModel: u})
	return u, v, list
}

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

func TestUserProfile(t *testing.T) {
	assertions := assert.New(t)
	u := &dal.User{ID: 1, Username: "test"}
	t.Run("success", func(t *testing.T) {
		userProfile := mockUserProfile(t, u, nil)
		resp, err := userProfile.GetProfile(context.TODO(), &types.UserInfoReq{UserID: u.ID})
		assertions.NoError(err)
		assertions.NotNil(resp)
		assertions.Equal(u.ID, resp.User.ID)
		assertions.Equal(u.Username, resp.User.Username)
	})
	t.Run("fail on FindOne", func(t *testing.T) {
		userProfile := mockUserProfile(t, nil, errors.New("error"))
		resp, err := userProfile.GetProfile(context.TODO(), &types.UserInfoReq{UserID: u.ID})
		assertions.Error(err)
		assertions.Nil(resp)
	})
}

func mockUserProfile(t *testing.T, u *dal.User, err error) *UserProfile {
	m := NewUserQuery(t)
	userProfile := NewUserProfile(UserProfileParam{UserModel: m})
	m.On("FindOne", mock.Anything, mock.Anything).Return(u, err)
	return userProfile
}

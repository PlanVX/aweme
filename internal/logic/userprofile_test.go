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

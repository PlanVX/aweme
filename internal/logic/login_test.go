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
	"github.com/PlanVX/aweme/internal/config"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestNewLogin(t *testing.T) {
	assertions := assert.New(t)
	text := []byte("a quick brown fox jumps over the lazy dog")
	bytearray, err := bcrypt.GenerateFromPassword(text, bcrypt.DefaultCost)
	assertions.NoError(err)
	arguments := &dal.User{Username: string(text), Password: bytearray}
	t.Run("success", func(t *testing.T) {
		login := mockLogin(t, arguments, nil)
		resp, err := login.Login(context.TODO(), &types.UserReq{Username: string(text), Password: string(text)})
		assertions.NoError(err)
		assertions.NotNil(resp)
		assertions.NotEmpty(resp.Token)
	})
	t.Run("fail on FindByUsername", func(t *testing.T) {
		login := mockLogin(t, nil, errors.New("error"))
		resp, err := login.Login(context.TODO(), &types.UserReq{Username: string(text), Password: string(text)})
		assertions.Error(err)
		assertions.Nil(resp)
	})
	t.Run("fail on bcrypt verify", func(t *testing.T) {
		login := mockLogin(t, arguments, nil)
		resp, err := login.Login(context.TODO(), &types.UserReq{Username: string(text), Password: "wrong password"})
		assertions.Error(err)
		assertions.Nil(resp)
	})
}

// mockLogin is a helper function to mock the login logic
// u is the user that will be returned by mocked UserQuery.FindByUsername method
// err is the error that will be returned by mocked UserQuery.FindByUsername method
func mockLogin(t *testing.T, u *dal.User, err error) *Login {
	m := NewUserQuery(t)
	login := NewLogin(LoginParam{UserQuery: m, J: mockJwt()})
	m.On("FindByUsername", mock.Anything, mock.Anything).Return(u, err)
	return login
}

// mockJwt is a helper function to mock the jwt signer
func mockJwt() *JWTSigner {
	c := config.Config{}
	c.JWT.Secret = "1234"
	c.JWT.TTL = 1234
	return NewJWTSigner(&c)
}

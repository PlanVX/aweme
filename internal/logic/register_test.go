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
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewRegister(t *testing.T) {
	assertions := assert.New(t)
	text := "a quick brown fox jumps over the lazy dog"
	userReq := types.UserReq{Username: text, Password: text}
	t.Run("success", func(t *testing.T) {
		register := mockRegister(t, nil)
		resp, err := register.Register(context.TODO(), &userReq)
		assertions.NoError(err)
		assertions.NotNil(resp)
		assertions.NotEmpty(resp.Token)
	})
	t.Run("fail on Insert", func(t *testing.T) {
		register := mockRegister(t, errors.New("error"))
		resp, err := register.Register(context.TODO(), &userReq)
		assertions.Error(err)
		assertions.Nil(resp)
	})
}

// mockRegister is a helper function to mock the UserQuery and return a Register
// err is the error that will be returned by the UserQuery.Insert method
func mockRegister(t *testing.T, err error) *Register {
	m := NewUserCommand(t)
	register := NewRegister(RegisterParam{UserCommand: m, J: mockJwt()})
	m.On("Insert", mock.Anything, mock.Anything).Return(err)
	return register
}

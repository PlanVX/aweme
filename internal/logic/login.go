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
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
)

type (
	// Login is the logic for login
	Login struct {
		userQuery dal.UserQuery
		signer    *JWTSigner
	}
	// LoginParam is the parameter for NewLogin
	LoginParam struct {
		fx.In
		UserQuery dal.UserQuery
		J         *JWTSigner
	}
)

// NewLogin is the constructor for Login
func NewLogin(param LoginParam) *Login {
	return &Login{
		userQuery: param.UserQuery,
		signer:    param.J,
	}
}

// Login 登陆逻辑
func (l *Login) Login(ctx context.Context, req *types.UserReq) (resp *types.UserResp, err error) {

	u, err := l.userQuery.FindByUsername(ctx, req.Username) // 根据用户名查找用户
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(req.Password)) // bcrypt 验证密码
	if err != nil {
		return nil, err
	}

	token, err := l.signer.genSignedToken(u.Username, u.ID)
	if err != nil {
		return nil, err
	}

	return &types.UserResp{UserID: u.ID, Token: token}, nil // 返回用户ID和token
}

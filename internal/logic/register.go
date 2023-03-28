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
	// Register is the logic for register
	Register struct {
		userCommand dal.UserCommand
		signer      *JWTSigner
	}
	// RegisterParam is the parameter for NewRegister
	RegisterParam struct {
		fx.In
		UserCommand dal.UserCommand
		J           *JWTSigner
	}
)

// NewRegister returns a new Register logic
func NewRegister(param RegisterParam) *Register {
	return &Register{
		userCommand: param.UserCommand,
		signer:      param.J,
	}
}

// Register 注册逻辑
// 注册账号，并把加密后的账号信息存入数据库，生成 token 返回
func (l *Register) Register(ctx context.Context, req *types.UserReq) (resp *types.UserResp, err error) {
	u := &dal.User{
		Username: req.Username,
	}

	// 使用 bcrypt 加密密码
	u.Password, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = l.userCommand.Insert(ctx, u) // 尝试保存到数据库
	if err != nil {
		return nil, err
	}

	token, err := l.signer.genSignedToken(u.Username, u.ID) // 生成 token
	if err != nil {
		return nil, err
	}

	return &types.UserResp{UserID: u.ID, Token: token}, nil
}

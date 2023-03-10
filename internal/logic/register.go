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
		userModel dal.UserModel
		signer    *JWTSigner
	}
	// RegisterParam is the parameter for NewRegister
	RegisterParam struct {
		fx.In
		UserModel dal.UserModel
		J         *JWTSigner
	}
)

// NewRegister returns a new Register logic
func NewRegister(param RegisterParam) *Register {
	return &Register{
		userModel: param.UserModel,
		signer:    param.J,
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

	err = l.userModel.Insert(ctx, u) // 尝试保存到数据库
	if err != nil {
		return nil, err
	}

	token, err := l.signer.genSignedToken(u.Username, u.ID) // 生成 token
	if err != nil {
		return nil, err
	}

	return &types.UserResp{UserID: u.ID, Token: token}, nil
}

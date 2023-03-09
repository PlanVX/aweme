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
		userModel dal.UserModel
		signer    *JWTSigner
	}
	// LoginParam is the parameter for NewLogin
	LoginParam struct {
		fx.In
		UserModel dal.UserModel
		J         *JWTSigner
	}
)

// NewLogin is the constructor for Login
func NewLogin(param LoginParam) *Login {
	return &Login{
		userModel: param.UserModel,
		signer:    param.J,
	}
}

// Login 登陆逻辑
func (l *Login) Login(ctx context.Context, req *types.UserReq) (resp *types.UserResp, err error) {

	u, err := l.userModel.FindByUsername(ctx, req.Username) // 根据用户名查找用户
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

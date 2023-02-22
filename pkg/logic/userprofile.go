package logic

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"go.uber.org/fx"
)

type (
	// UserProfile is the logic for querying user profile
	UserProfile struct {
		userModel dal.UserModel
	}
	// UserProfileParam is the parameter for NewUserProfile
	UserProfileParam struct {
		fx.In
		UserModel dal.UserModel
	}
)

// NewUserProfile returns a new UserProfile logic
func NewUserProfile(param UserProfileParam) *UserProfile {
	return &UserProfile{
		userModel: param.UserModel,
	}
}

// GetProfile 获取用户信息
// 根据用户 id 获取用户信息
func (u *UserProfile) GetProfile(ctx context.Context, req *types.UserInfoReq) (*types.UserInfoResp, error) {
	v, err := u.userModel.FindOne(ctx, req.UserID) // 根据用户 id 获取用户信息
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		User: covertUser(v),
	}, nil
}

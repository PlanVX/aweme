package api

import (
	"context"
	"github.com/PlanVX/aweme/pkg/types"
)

// NewRelationAction godoc
// @Summary 关系操作
// @Description 关系操作
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData types.RelationActionReq true "用户信息"
// @Success 200 {object} types.RelationActionResp
// @Router /relation/action/ [post]
func NewRelationAction() *Api {
	return &Api{
		Method: "POST",
		Path:   "/relation/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.RelationActionReq) (*types.RelationActionResp, error) {
			return nil, nil
		}),
	}
}

// NewRelationFollowList godoc
// @Summary 用户关注列表
// @Description 用户关注列表
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData types.RelationFollowListReq true "用户信息"
// @Success 200 {object} types.RelationFollowListResp
// @Router /relation/follow/list/ [get]
func NewRelationFollowList() *Api {
	return &Api{
		Method: "GET",
		Path:   "/relation/follow/list//",
		Handler: WrapperFunc(func(ctx context.Context, req *types.RelationFollowListReq) (*types.RelationFollowListResp, error) {
			return nil, nil
		}),
	}
}

// NewRelationFollowerList godoc
// @Summary 用户粉丝列表
// @Description 用户粉丝列表
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData types.RelationFollowerListReq true "用户信息"
// @Success 200 {object} types.RelationFollowerListResp
// @Router /relation/follower/list/ [get]
func NewRelationFollowerList() *Api {
	return &Api{
		Method: "GRT",
		Path:   "/relation/follower/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.RelationFollowerListReq) (*types.RelationFollowerListResp, error) {
			return nil, nil
		}),
	}
}

// NewRelationFriendList godoc
// @Summary 用户好友列表
// @Description 用户好友列表
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData types.RelationFriendListReq true "用户信息"
// @Success 200 {object} types.RelationFriendListResp
// @Router /relation/friend/list/ [get]
func NewRelationFriendList() *Api {
	return &Api{
		Method: "GET",
		Path:   "/relation/friend/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.RelationFriendListReq) (*types.RelationFriendListResp, error) {
			return nil, nil
		}),
	}
}

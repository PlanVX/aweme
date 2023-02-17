package api

import (
	"context"
	"github.com/PlanVX/aweme/pkg/types"
)

// NewFavoriteAction godoc
// @Summary 赞操作
// @Description 赞操作
// @Tags 互动接口
// @Produce json
// @Param favorite query types.FavoriteActionReq true "用户消息信息"
// @Success 200 {object} types.FavoriteActionResp
// @Router /favorite/action/ [get]
func NewFavoriteAction() *Api {
	return &Api{
		Method: "POST",
		Path:   "/favorite/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.FavoriteActionReq) (*types.FavoriteActionResp, error) {
			return nil, nil
		}),
	}
}

// NewFavoriteList godoc
// @Summary 收藏列表
// @Description 收藏列表
// @Tags 互动接口
// @Produce json
// @Param favorite query types.FavoriteListReq true "请求信息"
// @Success 200 {object} types.FavoriteListResp
// @Router /favorite/list/ [get]
func NewFavoriteList() *Api {
	return &Api{
		Method: "GET",
		Path:   "/favorite/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.FavoriteListReq) (*types.FavoriteListResp, error) {
			return nil, nil
		}),
	}
}

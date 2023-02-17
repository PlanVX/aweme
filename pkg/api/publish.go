package api

import (
	"context"
	"github.com/PlanVX/aweme/pkg/types"
)

// NewPublishList godoc
// @Summary 获取视频列表
// @Description 获取视频列表
// @Tags 基础接口
// @Produce json
// @Param param query types.PublishListReq true "获取视频列表参数"
// @Success 200 {object} types.PublishListResp
// @Router /publish/list/ [get]
func NewPublishList() *Api {
	return &Api{
		Method: "GET",
		Path:   "/publish/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.PublishListReq) (*types.PublishListResp, error) {
			return nil, nil
		}),
	}
}

package api

import (
	"github.com/PlanVX/aweme/internal/logic"
)

// NewPublishList godoc
// @Summary 获取视频列表
// @Description 获取视频列表
// @Tags 基础接口
// @Produce json
// @Param param query types.PublishListReq true "获取视频列表参数"
// @Success 200 {object} types.PublishListResp
// @Router /publish/list/ [get]
func NewPublishList(list *logic.PublishList) *Api {
	return &Api{
		Method:  "GET",
		Path:    "/publish/list/",
		Handler: WrapperFunc(list.PublishList),
	}
}

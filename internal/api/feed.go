package api

import (
	"github.com/PlanVX/aweme/internal/logic"
)

// NewFeed godoc
// @Summary 获取视频列表
// @Description 获取视频列表
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param feed query types.FeedReq true "视频列表请求参数"
// @Success 200 {object} types.FeedResp
// @Router /feed/ [get]
func NewFeed(f *logic.Feed) *Api {
	return &Api{
		Method:  "GET",
		Path:    "/feed/",
		Handler: WrapperFunc(f.Feed),
	}
}

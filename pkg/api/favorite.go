package api

import (
	"github.com/PlanVX/aweme/pkg/logic"
)

// NewFavoriteAction godoc
// @Summary 赞操作
// @Description 赞操作
// @Tags 互动接口
// @Produce json
// @Param favorite query types.FavoriteActionReq true "用户消息信息"
// @Success 200 {object} types.FavoriteActionResp
// @Router /favorite/action/ [get]
func NewFavoriteAction(like *logic.Like) *Api {
	return &Api{
		Method:  "POST",
		Path:    "/favorite/action/",
		Handler: WrapperFunc(like.Like),
	}
}

// NewFavoriteList godoc
// @Summary 点赞列表
// @Description 点赞列表
// @Tags 互动接口
// @Produce json
// @Param favorite query types.FavoriteListReq true "请求信息"
// @Success 200 {object} types.FavoriteListResp
// @Router /favorite/list/ [get]
func NewFavoriteList(list *logic.LikeList) *Api {
	return &Api{
		Method:  "GET",
		Path:    "/favorite/list/",
		Handler: WrapperFunc(list.LikeList),
	}
}

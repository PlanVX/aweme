package api

import "github.com/PlanVX/aweme/pkg/logic"

// NewUserInfo godoc
// @Summary 用户信息
// @Description 用户信息
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param user query types.UserInfoReq true "用户信息"
// @Success 200 {object} types.UserInfoResp
// @Router /user/info/ [get]
func NewUserInfo(profile *logic.UserProfile) *Api {
	return &Api{
		Path:    "/user/info/",
		Method:  "GET",
		Handler: WrapperFunc(profile.GetProfile),
	}
}

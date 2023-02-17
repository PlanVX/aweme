package api

import (
	"github.com/PlanVX/aweme/pkg/logic"
)

// NewRegister godoc
// @Summary 用户注册
// @Description 用户注册
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param user formData types.UserReq true "用户信息"
// @Success 200 {object} types.UserResp
// @Router /user/register/ [post]
func NewRegister(l *logic.Register) *Api {
	return &Api{
		Method:  "POST",
		Path:    "/user/register/",
		Handler: WrapperFunc(l.Register),
	}
}

// NewLogin godoc
// @Summary 用户登陆
// @Description 用户登陆
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param user formData types.UserReq true "用户信息"
// @Success 200 {object} types.UserResp
// @Router /user/login/ [post]
func NewLogin(l *logic.Login) *Api {
	return &Api{
		Method:  "POST",
		Path:    "/user/login/",
		Handler: WrapperFunc(l.Login),
	}
}

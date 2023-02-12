package api

import "context"

type UserResp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

type UserReq struct {
	Username string `query:"username" json:"username" query:"username"` // 用户名
	Password string `query:"password" json:"password" query:"password"` // 密码
}

type RegisterApiParam struct{}

// NewRegisterApi godoc
// @Summary 用户注册
// @Description 用户注册
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param user formData UserReq true "用户信息"
// @Success 200 {object} UserResp
// @Router /user/register [post]
func NewRegisterApi(param RegisterApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/user/register/",
		Handler: WrapperFunc(func(ctx context.Context, req *UserReq) (*UserResp, error) {
			return nil, nil
		}),
	}
}

type LoginApiParam struct{}

// NewLoginApi godoc
// @Summary 用户登陆
// @Description 用户登陆
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param user formData UserReq true "用户信息"
// @Success 200 {object} UserResp
// @Router /user/login [post]
func NewLoginApi(param LoginApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/user/login/",
		Handler: WrapperFunc(func(ctx context.Context, req *UserReq) (*UserResp, error) {
			return nil, nil
		}),
	}
}

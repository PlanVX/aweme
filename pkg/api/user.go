package api

import "context"

type UserGetReq struct {
	Token  string `query:"token" json:"token"`     // 用户鉴权token
	UserID int64  `query:"user_id" json:"user_id"` // 用户id
}

type UserGetResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	User       User   `json:"user"`        // 用户信息
}

type UserApiParam struct{}

// NewUserApi godoc
// @Summary 获取登录用户信息
// @Description 获取登录用户信息
// @Tags 用户接口
// @Accept json
// @Produce  json
// @Param user_id query int64 true "用户id"
// @Param token query string true "用户鉴权token"
// @Success 200 {object} UserResp
// @Router /user [get]
func NewUserApi(param UserApiParam) *Api {
	return &Api{
		Method: "GET",
		Path:   "/user/",
		Handler: WrapperFunc(func(ctx context.Context, req *UserGetReq) (*UserResp, error) {
			// 获取用户信息
			return nil, nil
		}),
	}
}

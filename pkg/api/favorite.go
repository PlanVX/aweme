package api

import "context"

type FavoriteActionReq struct {
	Token      string `query:"token" json:"token"`             // 用户鉴权token
	VideoID    int64  `query:"video_id" json:"video_id"`       // 视频id
	ActionType int32  `query:"action_type" json:"action_type"` // 1-点赞，2-取消点赞
}

type FavoriteActionResp struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

type FavoriteActionApiParam struct{}

// NewFavoriteActionApi godoc
// @Summary 赞操作
// @Description 赞操作
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param favorite query FavoriteActionReq true "用户视频信息"
// @Success 200 {object} FavoriteActionResp
// @Router /favorite/action [get]
func NewFavoriteActionApi(param FavoriteActionApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/favorite/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *FavoriteActionReq) (*FavoriteActionResp, error) {
			return nil, nil
		}),
	}
}

// 用户点赞列表请求
type FavoriteListReq struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// 用户点赞列表响应
type FavoriteListResp struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户点赞视频列表
}

type NewFavoriteListApiParam struct{}

// NewFavoriteListApi godoc
// @Summary 收藏列表
// @Description 收藏列表
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param favorite query string true "用户信息"
// @Success 200 {object} FavoriteListResp
// @Router /favorite/list [get]
func NewFavoriteListApi(param NewFavoriteListApiParam) *Api {
	return &Api{
		Method: "GET",
		Path:   "/favorite/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *FavoriteListReq) (*FavoriteListResp, error) {
			return nil, nil
		}),
	}
}

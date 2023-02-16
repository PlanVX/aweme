package api

import "context"

type PublishActionReq struct {
	Token string `query:"token" json:"token"` // 用户鉴权token
	Data  []byte `query:"data" json:"data"`   // 视频数据
	Title string `query:"title" json:"title"` // 视频标题
}

type PublishActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type PublishListReq struct {
	UserReq
}

type PublishListResp struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户发布的视频列表
}

type PublishActionApiParam struct{}

type PublishListApiParam struct{}

// NewPublishActionApi godoc
// @Summary 视频发布
// @Description 用户发布短视频
// @Tags 视频接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param param body PublishActionReq true "发布视频参数"
// @Success 200 {object} PublishActionResp
// @Router /publish/action/ [post]
func NewPublishActionApi(param PublishActionApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/publish/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *PublishActionReq) (*PublishActionResp, error) {
			return nil, nil
		}),
	}
}

// NewPublishListApi godoc
// @Summary 获取视频列表
// @Description 获取视频列表
// @Tags 视频接口
// @Accept json
// @Produce json
// @Param param body PublishListReq true "获取视频列表参数"
// @Success 200 {object} PublishListResp
// @Router /publish/list [get]
func NewPublishListApi(param PublishListApiParam) *Api {
	return &Api{
		Method: "GET",
		Path:   "/publish/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *PublishListReq) (*PublishListResp, error) {
			return nil, nil
		}),
	}
}

package api

import "context"

type PublishReq struct {
	Token string `query:"token" json:"token"` // 用户鉴权token
	Data  []byte `query:"data" json:"data"`   // 视频数据
	Title string `query:"title" json:"title"` // 视频标题
}

type PublishResp struct {
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

type PublishActionApiParam struct {
	Token string `json:"token"` // 用户鉴权token
	Data  []byte `json:"data"`  // 视频数据
	Title string `json:"title"` // 视频标题
}

type PublishListApiParam struct {
	UserID int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// NewPublishActionApi returns a new video publishing API instance.
func NewPublishActionApi(param PublishActionApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/video/publish/",
		Handler: WrapperFunc(func(ctx context.Context, req *PublishReq) (*PublishResp, error) {
			return nil, nil
		}),
	}
}

// NewPublishListApi returns a new video list API instance.
func NewPublishListApi(param PublishListApiParam) *Api {
	return &Api{
		Method: "GET",
		Path:   "/video/list/",
		Handler: WrapperFunc(func(ctx context.Context, req *PublishListReq) (*PublishListResp, error) {
			return nil, nil
		}),
	}
}

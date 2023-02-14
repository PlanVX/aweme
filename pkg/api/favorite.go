package api

import "context"

type FavoriteActionRequest struct {
	Token      string `query:"token" json:"token"`             // 用户鉴权token
	VideoID    int64  `query:"video_id" json:"video_id"`       // 视频id
	ActionType int32  `query:"action_type" json:"action_type"` // 1-点赞，2-取消点赞
}

type FavoriteActionResponse struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

type FavoriteActionApiParam struct{}

// NewFavoriteActionApi returns a new video favorite action API instance.
func NewFavoriteActionApi(param FavoriteActionApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/video/favorite/",
		Handler: WrapperFunc(func(ctx context.Context, req *FavoriteActionRequest) (*FavoriteActionResponse, error) {
			return nil, nil
		}),
	}
}

// 用户点赞列表请求
type FavoriteListReq struct {
	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

// 用户点赞列表响应
type FavoriteListResp struct {
	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`
	VideoList  []Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list,omitempty"`
}

type NewFavoriteListApiParam struct{}

// NewFavoriteListApi godoc
// @Summary 收藏列表
// @Description 收藏列表
// @Tags 收藏接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user_id query string true "用户ID"
// @Success 200 {object} FavoriteListResp
// @Router /favorite/list [get]
func NewFavoriteListApi(param NewFavoriteListApiParam) *Api {
	return &Api{
		Method: "GET",
		Path:   "/favorite/list",
		Handler: WrapperFunc(func(ctx context.Context, req *FavoriteListReq) (*FavoriteListResp, error) {
			return nil, nil
		}),
	}
}

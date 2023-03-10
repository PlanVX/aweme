package types

// PublishListReq is the request of publish list api
type PublishListReq struct {
	Token  string `query:"token" form:"token" json:"token"`       // token
	UserID int64  `query:"user_id" form:"user_id" json:"user_id"` // user id
}

// PublishListResp is the response of publish list api
type PublishListResp struct {
	Response
	VideoList []*Video `json:"video_list"` // 用户发布的视频列表
}

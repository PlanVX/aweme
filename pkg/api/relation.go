package api

import "context"

// RelationFollowListReq 用于获取关注列表请求
type RelationFollowListReq struct {
	UserID int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// RelationFollowListResp 用于获取关注列表响应
type RelationFollowListResp struct {
	StatusCode int32      `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string     `json:"status_msg"`  // 返回状态描述
	UserList   []UserResp `json:"user_list"`   // 用户信息列表
}

// RelationFollowerListReq 用于获取粉丝列表请求
type RelationFollowerListReq struct {
	UserID int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// RelationFollowerListResp 用于获取粉丝列表响应
type RelationFollowerListResp struct {
	StatusCode int32      `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string     `json:"status_msg"`  // 返回状态描述
	UserList   []UserResp `json:"user_list"`   // 用户信息列表
}

// RelationFriendListReq 用于获取好友列表请求
type RelationFriendListReq struct {
	UserID int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// RelationFriendListResp 用于获取好友列表响应
type RelationFriendListResp struct {
	StatusCode int32            `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string           `json:"status_msg"`  // 返回状态描述
	UserList   []FriendUserResp `json:"user_list"`   // 用户信息列表
}

// FriendUserResp 好友用户信息
type FriendUserResp struct {
	UserResp
	Message string `json:"message,omitempty"` // 和该好友的最新聊天消息
	MsgType int64  `json:"msg_type"`          // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

// RelationActionReq 关系操作请求
type RelationActionReq struct {
	Token      string `json:"token"`       // 用户鉴权 token
	ToUserID   int64  `json:"to_user_id"`  // 对方用户 id
	ActionType int32  `json:"action_type"` // 1-关注，2-取消关注
}

// RelationActionResp 关系操作响应
type RelationActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type RelationActionApiParam struct{}

// NewRelationActionApi godoc
// @Summary 关系操作
// @Description 关系操作
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData RelationActionReq true "用户信息"
// @Success 200 {object} RelationActionResp
// @Router /relation/action/ [post]
func NewRelationActionApi(param RelationActionApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/relation/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *RelationActionReq) (*RelationActionResp, error) {
			return nil, nil
		}),
	}
}

type RelationFollowListApiParam struct{}

// NewRelationFollowListApi godoc
// @Summary 用户关注列表
// @Description 用户关注列表
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData RelationFollowListReq true "用户信息"
// @Success 200 {object} RelationFollowListResp
// @Router /relation/follow_list/ [get]
func NewRelationFollowListApi(param RelationFollowListApiParam) *Api {
	return &Api{
		Method: "GET",
		Path:   "/relation/follow_list/",
		Handler: WrapperFunc(func(ctx context.Context, req *RelationFollowListReq) (*RelationFollowListResp, error) {
			return nil, nil
		}),
	}
}

type RelationFollowerListApiParam struct{}

// NewRelationFollowerListApi godoc
// @Summary 用户粉丝列表
// @Description 用户粉丝列表
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData RelationFollowerListReq true "用户信息"
// @Success 200 {object} RelationFollowerListResp
// @Router /relation/follower_list [get]
func NewRelationFollowerListApi(param RelationFollowerListApiParam) *Api {
	return &Api{
		Method: "GRT",
		Path:   "/relation/follower_list/",
		Handler: WrapperFunc(func(ctx context.Context, req *RelationFollowerListReq) (*RelationFollowerListResp, error) {
			return nil, nil
		}),
	}
}

type RelationFriendListApiParam struct{}

// NewRelationFriendListApi godoc
// @Summary 用户好友列表
// @Description 用户好友列表
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param relation formData RelationFriendListReq true "用户信息"
// @Success 200 {object} RelationFriendListResp
// @Router /relation/friend_list [get]
func NewRelationFriendListApi(param RelationFriendListApiParam) *Api {
	return &Api{
		Method: "GET",
		Path:   "/relation/friend_list/",
		Handler: WrapperFunc(func(ctx context.Context, req *RelationFriendListReq) (*RelationFriendListResp, error) {
			return nil, nil
		}),
	}
}

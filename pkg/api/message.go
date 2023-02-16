package api

import "context"

type MessageChatReq struct {
	Token    string `json:"token" query:"token"`           // 用户鉴权token
	ToUserID int64  `json:"to_user_id" query:"to_user_id"` // 对方用户id
}

type MessageChatResp struct {
	StatusCode  int32     `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg"`   // 返回状态描述
	MessageList []Message `json:"message_list"` // 消息列表
}

type MessageActionReq struct {
	Token      string `json:"token" query:"token"`             // 用户鉴权token
	ToUserID   int64  `json:"to_user_id" query:"to_user_id"`   // 对方用户id
	ActionType int32  `json:"action_type" query:"action_type"` // 1-发送消息
	Content    string `json:"content" query:"content"`         // 消息内容
}

type MessageActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type MessageChatApiParam struct{}

// NewMessageChatApi godoc
// @Summary 聊天记录
// @Description 聊天记录
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param message query MessageChatReq true "用户信息"
// @Success 200 {object} MessageChatResp
// @Router /message/chat [get]
func NewMessageChatApi(param MessageChatApiParam) *Api {
	return &Api{
		Method: "GET",
		Path:   "/message/chat/",
		Handler: WrapperFunc(func(ctx context.Context, req *MessageChatReq) (*MessageChatResp, error) {
			return nil, nil
		}),
	}
}

type MessageActionApiParam struct{}

// NewMessageActionApi godoc
// @Summary 消息操作
// @Description 消息操作
// @Tags 基础接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param message query MessageActionReq true "用户消息信息"
// @Success 200 {object} MessageActionResp
// @Router /message/action [post]
func NewMessageActionApi(param MessageActionApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/message/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *MessageActionReq) (*MessageActionResp, error) {
			return nil, nil
		}),
	}
}

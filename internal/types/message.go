package types

// MessageListReq chat message list request
type MessageListReq struct {
	Token    string `json:"token" query:"token"`           // 用户鉴权token
	ToUserID int64  `json:"to_user_id" query:"to_user_id"` // 对方用户id
}

// MessageListResp chat message list response
type MessageListResp struct {
	StatusCode  int32      `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string     `json:"status_msg"`   // 返回状态描述
	MessageList []*Message `json:"message_list"` // 消息列表
}

// MessageActionReq chat action request
type MessageActionReq struct {
	Token      string `json:"token" query:"token"`             // 用户鉴权token
	ToUserID   int64  `json:"to_user_id" query:"to_user_id"`   // 对方用户id
	ActionType int32  `json:"action_type" query:"action_type"` // 1-发送消息
	Content    string `json:"content" query:"content"`         // 消息内容
}

// MessageActionResp chat action response
type MessageActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

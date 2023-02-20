package types

// CommentActionReq 评论操作请求
type CommentActionReq struct {
	Token       string `query:"token" form:"token"`               // 用户鉴权token
	VideoID     int64  `query:"video_id" form:"video_id"`         // 视频id
	ActionType  int32  `query:"action_type" form:"action_type"`   // 1-发布评论，2-删除评论
	CommentText string `query:"comment_text" form:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentID   int64  `query:"comment_id" form:"comment_id"`     // 要删除的评论id，在action_type=2的时候使用
}

// CommentActionResp 评论操作响应
type CommentActionResp struct {
	StatusCode int32    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `json:"status_msg"`  // 返回状态描述
	Comment    *Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
}

// CommentListReq 评论列表请求
type CommentListReq struct {
	Token   string `json:"token" query:"token" form:"token"`          // 用户鉴权token
	VideoID int64  `json:"video_id" query:"video_id" form:"video_id"` // 视频id
}

// CommentListResp 评论列表响应
type CommentListResp struct {
	StatusCode  int32      `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string     `json:"status_msg"`   // 返回状态描述
	CommentList []*Comment `json:"comment_list"` // 评论列表
}

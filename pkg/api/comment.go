package api

import "context"

// 评论操作请求
type CommentActionReq struct {
	Token       string `json:"token"`        // 用户鉴权token
	VideoID     int64  `json:"video_id"`     // 视频id
	ActionType  int32  `json:"action_type"`  // 1-发布评论，2-删除评论
	CommentText string `json:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentID   int64  `json:"comment_id"`   // 要删除的评论id，在action_type=2的时候使用
}

// 评论操作响应
type CommentActionResp struct {
	StatusCode int32    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `json:"status_msg"`  // 返回状态描述
	Comment    *Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
}

// 评论列表请求
type CommentListReq struct {
	Token   string `json:"token"`    // 用户鉴权token
	VideoID int64  `json:"video_id"` // 视频id
}

// 评论列表响应
type CommentListResp struct {
	StatusCode  int32      `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string     `json:"status_msg"`   // 返回状态描述
	CommentList []*Comment `json:"comment_list"` // 评论列表
}

type NewCommentActionApiParam struct {
	CommentId string `json:"comment_id" form:"comment_id" binding:"required"`
	Action    string `json:"action" form:"action" binding:"required"`
}

type NewCommentListApiParam struct {
	UserId   string `json:"user_id" form:"user_id" binding:"required"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type CommentActionApiParam struct{}

// NewCommentActionApi godoc
// @Summary 评论操作
// @Description 评论操作
// @Tags 评论接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user formData CommentActionReq true "评论信息"
// @Success 200 {object} CommentActionResp
// @Router /comment/action [post]
func NewCommentActionApi(param CommentActionApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/comment/action",
		Handler: WrapperFunc(func(ctx context.Context, req *CommentActionReq) (*CommentActionResp, error) {
			return nil, nil
		}),
	}
}

type CommentListApi struct{}

// NewCommentListApi godoc
// @Summary 评论列表
// @Description 评论列表
// @Tags 评论接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user_id query string true "用户ID"
// @Success 200 {object} CommentListResp
// @Router /comment/list [get]
func NewCommentListApi(param CommentListApi) *Api {
	return &Api{
		Method: "GET",
		Path:   "/comment/list",
		Handler: WrapperFunc(func(ctx context.Context, req *CommentListReq) (*CommentListResp, error) {
			return nil, nil
		}),
	}
}

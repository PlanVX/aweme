package api

import (
	"context"
	"github.com/PlanVX/aweme/pkg/types"
)

// NewMessageChat godoc
// @Summary 聊天记录
// @Description 聊天记录
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param message query types.MessageListReq true "用户信息"
// @Success 200 {object} types.MessageListResp
// @Router /message/chat/ [get]
func NewMessageChat() *Api {
	return &Api{
		Method: "GET",
		Path:   "/message/chat/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.MessageListReq) (*types.MessageListResp, error) {
			return nil, nil
		}),
	}
}

// NewMessageAction godoc
// @Summary 消息操作
// @Description 消息操作
// @Tags 社交接口
// @Accept x-www-form-urlencoded
// @Produce json
// @Param message query types.MessageActionReq true "用户消息"
// @Success 200 {object} types.MessageActionResp
// @Router /message/action/ [post]
func NewMessageAction() *Api {
	return &Api{
		Method: "POST",
		Path:   "/message/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *types.MessageActionReq) (*types.MessageActionResp, error) {
			return nil, nil
		}),
	}
}

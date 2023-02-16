package logic

import (
	"context"
	"github.com/PlanVX/aweme/pkg/dal"
	"github.com/PlanVX/aweme/pkg/types"
	"go.uber.org/fx"
)

type (
	// Upload is the logic for upload video
	Upload struct {
		videoModel dal.VideoModel
	}
	// UploadParam is the param for NewUpload
	UploadParam struct {
		fx.In
		VideoModel dal.VideoModel
	}
)

// NewUpload returns a new Upload logic
func NewUpload(param UploadParam) *Upload {
	return &Upload{
		videoModel: param.VideoModel,
	}
}

// UploadVideo publishes a video
func (u *Upload) UploadVideo(_ context.Context, _ *types.UploadReq) (*types.UploadResp, error) {
	return &types.UploadResp{}, nil
}

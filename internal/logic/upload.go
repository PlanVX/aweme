/*
 * Copyright (c) 2023 The PlanVX Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logic

import (
	"context"
	"github.com/PlanVX/aweme/internal/dal"
	"github.com/PlanVX/aweme/internal/types"
	"go.uber.org/fx"
	"strconv"
	"time"
)

type (
	// Upload is the logic for upload video
	Upload struct {
		videoCommand dal.VideoCommand
		uploader     Uploader
	}
	// UploadParam is the param for NewUpload
	UploadParam struct {
		fx.In
		VideoCommand dal.VideoCommand
		Helper       Uploader
	}
)

// NewUpload returns a new Upload logic
func NewUpload(param UploadParam) *Upload {
	return &Upload{
		videoCommand: param.VideoCommand,
		uploader:     param.Helper,
	}
}

// UploadVideo publishes a video
func (u *Upload) UploadVideo(c context.Context, req *types.UploadReq) (*types.UploadResp, error) {
	owner, _ := c.Value(ContextKey).(int64) // get the owner from context

	key := strconv.FormatInt(time.Now().UnixNano(), 10) + req.FileName // generate a unique key for the video

	upload, err := u.uploader.Upload(c, UploadInput{Key: key, Value: req.Data}) // upload the video
	if err != nil {
		return nil, err
	}

	// insert the video into database
	err = u.videoCommand.Insert(c, &dal.Video{VideoURL: upload, UserID: owner, Title: req.Title})
	if err != nil {
		return nil, err
	}

	return &types.UploadResp{}, nil
}

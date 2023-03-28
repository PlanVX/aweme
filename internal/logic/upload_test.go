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
	"bytes"
	"context"
	"errors"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUpload(t *testing.T) {
	assertions := assert.New(t)
	url := "url"

	buffer := bytes.NewBuffer([]byte("test"))

	req := &types.UploadReq{Data: buffer}
	t.Run("success", func(t *testing.T) {
		uploader, v, upload := mockUploadLogic(t)
		v.On("Insert", mock.Anything, mock.Anything).Return(nil)
		uploader.On("Upload", mock.Anything, mock.Anything).Return(url, nil)
		resp, err := upload.UploadVideo(context.TODO(), req)
		assertions.NoError(err)
		assertions.NotNil(resp)
	})
	t.Run("fail on Upload", func(t *testing.T) {
		uploader, _, upload := mockUploadLogic(t)
		uploader.On("Upload", mock.Anything, mock.Anything).Return("", errors.New("error"))
		resp, err := upload.UploadVideo(context.TODO(), req)
		assertions.Error(err)
		assertions.Nil(resp)
	})
	t.Run("fail on Insert", func(t *testing.T) {
		uploader, v, upload := mockUploadLogic(t)
		uploader.On("Upload", mock.Anything, mock.Anything).Return(url, nil)
		v.On("Insert", mock.Anything, mock.Anything).Return(errors.New("INSERT ERROR"))
		resp, err := upload.UploadVideo(context.TODO(), req)
		assertions.Error(err)
		assertions.Nil(resp)
	})
}

func mockUploadLogic(t *testing.T) (*MockUploader, *VideoCommand, *Upload) {
	uploader := NewMockUploader(t)
	v := NewVideoCommand(t)
	upload := NewUpload(UploadParam{VideoCommand: v, Helper: uploader})
	return uploader, v, upload
}

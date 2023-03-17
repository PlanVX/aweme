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

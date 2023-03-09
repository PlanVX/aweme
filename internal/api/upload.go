package api

import (
	"github.com/PlanVX/aweme/internal/logic"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/labstack/echo/v4"
)

// NewUpload upload video
// @Summary 上传视频
// @Description 上传视频
// @Tags 基础接口
// @Produce json
// @Accept mpfd
// @Param token formData string true "Authorization token"
// @Param title formData string true "Data title"
// @Param data formData file true "Data"
// @Success 200 {object} types.UploadResp
// @Router /publish/action/ [post]
func NewUpload(upload *logic.Upload) *Api {
	return &Api{
		Method: "POST",
		Path:   "/publish/action/",
		Handler: func(c echo.Context) error {
			title := c.FormValue("title") // title param from form
			if title == "" {
				return echo.ErrBadRequest
			}
			file, err := c.FormFile("data") // file from form
			if err != nil {
				return err
			}
			f, err := file.Open()
			if err != nil {
				return err
			}
			resp, err := upload.UploadVideo(c.Request().Context(),
				&types.UploadReq{Title: title, Data: f, FileName: file.Filename}) // handle upload
			if err != nil {
				return err
			}
			return c.JSON(200, resp)
		},
	}
}

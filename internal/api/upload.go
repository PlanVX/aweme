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
func NewUpload(upload *logic.Upload) *API {
	return &API{
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

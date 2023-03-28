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
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

// WrapperFunc wrapper function to echo.HandlerFunc
func WrapperFunc[Req any, Resp any](biz func(context.Context, *Req) (*Resp, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(Req)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "unsupported request parameters")
		}
		if err := c.Validate(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid request parameters")
		}
		resp, err := biz(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(200, resp)
	}
}

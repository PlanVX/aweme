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
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestReq struct {
	In string `form:"in" json:"in" validate:"required,alphanum,min=1,max=16" query:"in"`
}

type TestResp struct {
	Out string `json:"out"`
}

func Biz(_ context.Context, req *TestReq) (*TestResp, error) {
	if req.In == "error" {
		return nil, echo.NewHTTPError(500, "error")
	}
	resp := new(TestResp)
	resp.Out = req.In
	return resp, nil
}

func TestWrapperFunc(t *testing.T) {
	handler := WrapperFunc(Biz)
	e := echo.New()
	e.Validator = NewCustomValidator()
	var testCases = []struct {
		jsonStr string
		wantErr bool
		resp    string
	}{
		{`{"in":"hello"}`, false, `{"out":"hello"}`},
		{`{f}`, true, ``},
		{`{"in":"error"}`, true, ``},
		{`{"in":"error"}`, true, ``},
		{`{"in":"error."}`, true, ``},
	}
	for _, tc := range testCases {
		t.Run(tc.jsonStr, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/", strings.NewReader(tc.jsonStr))
			request.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()
			c := e.NewContext(request, recorder)
			err := handler(c)
			cond := require.New(t)
			if tc.wantErr {
				cond.Error(err)
				return
			}
			cond.NoError(err)
			cond.JSONEq(tc.resp, recorder.Body.String())
		})

	}
}

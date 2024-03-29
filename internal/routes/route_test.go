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

package routes

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/PlanVX/aweme/internal/api"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/PlanVX/aweme/internal/logic"
	"github.com/PlanVX/aweme/internal/types"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"net/http/httptest"
	"testing"
)

func TestModule(t *testing.T) {
	app := fxtest.New(t, fx.Provide(
		func() *config.Config {
			c := &config.Config{}
			c.API.Prefix = "/api"
			c.API.Address = ":8421"
			return c
		},
		logic.NewJWTSigner,
		zap.NewDevelopment,
		fx.Annotate(mockApis, fx.ResultTags(`group:"public"`, `group:"optional"`, `group:"private"`)),
	), Module,
	)
	assert.NoError(t, app.Start(context.Background()))
	assert.NoError(t, app.Stop(context.Background()))
}

// mockApis is a mock for providing three types of apis
func mockApis() (*api.API, *api.API, *api.API) {
	apis := &api.API{Path: "/public", Method: "GET", Handler: func(c echo.Context) error {
		return c.JSON(200, "mock")
	}}
	return apis, apis, apis
}

func TestNewEcho(t *testing.T) {
	example := zap.NewExample()
	e := NewEcho(example)
	assertions := assert.New(t)
	assertions.NotNil(e)
	e.Add("GET", "/test", func(c echo.Context) error {
		param := c.QueryParam("param")
		switch param {
		case "echo": // echo error
			return echo.NewHTTPError(400, "echo error")
		case "panic": // panic error
			panic("panic error")
		case "normal": // normal error
			return errors.New("normal error")
		default: // normal response
			return c.JSON(200, "ok")
		}
	})

	resp := types.Response{
		Code: 500,
		Msg:  "failed",
	}
	var testCases = []struct {
		name     string
		param    string
		httpCode int
		resp     types.Response
	}{{
		name:     "echo error",
		param:    "echo",
		httpCode: 200,
		resp: types.Response{
			Code: 400,
			Msg:  "echo error",
		},
	}, {
		name:     "panic error",
		param:    "panic",
		httpCode: 200,
		resp:     resp,
	},
		{
			name:     "normal error",
			param:    "normal",
			httpCode: 200,
			resp:     resp,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test?param="+tc.param, nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			assertions.Equal(tc.httpCode, rec.Code)
			var resp types.Response
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assertions.NoError(err)
			assertions.Equal(tc.resp, resp)
		})
	}
}

func Test_prefixSkipper(t *testing.T) {
	skipper := prefixSkipper("/api")
	var testCases = []struct {
		name string
		path string
		want bool
	}{
		{
			name: " starts with /api",
			path: "/api/test",
			want: false,
		},
		{
			name: "not starts with /api",
			path: "/test",
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.path, nil)
			resp := httptest.NewRecorder()
			e := echo.New()
			newContext := e.NewContext(req, resp)
			assert.Equal(t, tc.want, skipper(newContext))
		})
	}

}

type Request struct {
	Data string `query:"data"`
}

func TestBind(t *testing.T) {
	request := httptest.NewRequest("POST", "/test?data=test", nil)
	e := echo.New()
	newContext := e.NewContext(request, nil)
	v := &Request{}
	err := NewCustomBinder().Bind(v, newContext)
	assert.NoError(t, err)
	assert.Equal(t, "test", v.Data)
}

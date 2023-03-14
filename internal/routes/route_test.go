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
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
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
		tracesdk.NewTracerProvider,
		fx.Annotate(mockApis, fx.ResultTags(`group:"public"`, `group:"optional"`, `group:"private"`)),
	), Module,
	)
	assert.NoError(t, app.Start(context.Background()))
	assert.NoError(t, app.Stop(context.Background()))
}

// mockApis is a mock for providing three types of apis
func mockApis() (*api.Api, *api.Api, *api.Api) {
	apis := &api.Api{Path: "/public", Method: "GET", Handler: func(c echo.Context) error {
		return c.JSON(200, "mock")
	}}
	return apis, apis, apis
}

func TestNewEcho(t *testing.T) {
	example := zap.NewExample()
	provider := tracesdk.NewTracerProvider()
	e := NewEcho(example, provider)
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

package routes

import (
	"context"
	"github.com/PlanVX/aweme/pkg/api"
	"github.com/PlanVX/aweme/pkg/config"
	"github.com/PlanVX/aweme/pkg/logic"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
	"testing"
)

func TestModule(t *testing.T) {
	app := fxtest.New(t, fx.Provide(
		func() *config.Config { return &config.Config{} },
		logic.NewJWTSigner,
		zap.NewDevelopment,
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

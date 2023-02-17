package main

import (
	"github.com/PlanVX/aweme/pkg/api"
	"github.com/PlanVX/aweme/pkg/config"
	"github.com/PlanVX/aweme/pkg/dal/query"
	"github.com/PlanVX/aweme/pkg/logic"
	"github.com/PlanVX/aweme/pkg/routes"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// @title aweme
// @version 1.0
// @description aweme api
// @contact.name PlanVX
// @contact.url https://github.com/PlanVX
// @license.name Apache 2.0
// @license.url https://github.com/PlanVX/aweme/blob/main/LICENSE
// @host localhost:8080
// @BasePath /v1
func main() {
	app := fx.New(
		fx.Provide(config.NewConfig, newZapLogger),
		fx.WithLogger(fxLogger),
		query.Module,
		logic.Module,
		api.Module,
		routes.Module,
	)
	app.Run()
}

// replace the default logger with wrapped zap logger
func fxLogger(logger *zap.Logger) fxevent.Logger { return &fxevent.ZapLogger{Logger: logger} }

// newZapLogger returns a new zap logger
func newZapLogger(c *config.Config) (*zap.Logger, error) {
	if !c.Release {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}

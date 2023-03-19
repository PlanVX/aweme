//go:build swagger
// +build swagger

package routes

import (
	"github.com/PlanVX/aweme/docs"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// trick to make sure decorators is initialized before it is used
var _ = func() any {
	decorators = append(decorators, registerSwagger)
	return nil
}()

// registerSwagger registers swagger docs route
func registerSwagger(config *config.Config, e *echo.Echo) *echo.Echo {
	docs.SwaggerInfo.BasePath = config.API.Prefix // set swagger base path same as echo group prefix
	e.GET("/swagger/*", echoSwagger.WrapHandler)  // add swagger docs route
	return e
}

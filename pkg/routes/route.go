package routes

import (
	"context"
	"github.com/PlanVX/aweme/docs"
	"github.com/PlanVX/aweme/pkg/api"
	"github.com/PlanVX/aweme/pkg/config"
	"github.com/PlanVX/aweme/pkg/logic"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

// NewEcho returns a new echo instance and basic middleware is added
func NewEcho(logger *zap.Logger) *echo.Echo {
	e := echo.New()
	e.HideBanner = true              // hide echo banner
	e.HidePort = true                // hide port in log
	e.Use(echozap.ZapLogger(logger)) // use zap logger to replace default logger
	// add recover middleware so when panic happens, it will be recovered to centralize error handling
	e.Use(middleware.Recover())
	return e
}

// AddRoutersParam is the param for AddRouters
type AddRoutersParam struct {
	fx.In
	PublicApis   []*api.Api `group:"public"`
	OptionalApis []*api.Api `group:"optional"`
	PrivateApis  []*api.Api `group:"private"`
	E            *echo.Echo
	Signer       *logic.JWTSigner
	Config       *config.Config
}

// AddRouters adds all the routes to echo
func AddRouters(param AddRoutersParam) *echo.Echo {
	prefix := param.Config.API.Prefix // get api prefix from config
	// group apis with common prefix
	group := param.E.Group(prefix)
	for _, h := range param.PublicApis { // add public apis
		group.Add(h.Method, h.Path, h.Handler)
	}
	// 写入白名单
	param.Signer.AddWhiteList(lo.Map(param.OptionalApis, func(h *api.Api, _ int) string { return prefix + h.Path }))
	group.Use(param.Signer.NewJWTMiddleware()) // use jwt middleware
	for _, h := range param.OptionalApis {     // add optional apis
		group.Add(h.Method, h.Path, h.Handler)
	}
	for _, h := range param.PrivateApis { // add private apis
		group.Add(h.Method, h.Path, h.Handler)
	}

	docs.SwaggerInfo.BasePath = prefix                 // set swagger base path same as echo group prefix
	param.E.GET("/swagger/*", echoSwagger.WrapHandler) // add swagger docs route
	return param.E
}

// StartServer starts the HTTP server in fx.Lifecycle, so that it can be gracefully shutdown.
func StartServer(lf fx.Lifecycle, e *echo.Echo) {
	server := &http.Server{Addr: ":8080", Handler: e}
	lf.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		}()
		return nil
	}, OnStop: func(ctx context.Context) error { return server.Shutdown(ctx) }})
}

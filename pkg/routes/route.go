package routes

import (
	"context"
	"github.com/PlanVX/aweme/docs"
	"github.com/PlanVX/aweme/pkg/api"
	"github.com/PlanVX/aweme/pkg/config"
	"github.com/PlanVX/aweme/pkg/logic"
	"github.com/PlanVX/aweme/pkg/types"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// CustomBinder is a custom binder to bind request body to struct
type CustomBinder struct {
	d echo.DefaultBinder
}

// NewCustomBinder returns a new CustomBinder
func NewCustomBinder() *CustomBinder {
	return &CustomBinder{d: echo.DefaultBinder{}}
}

// Bind implements echo.Binder interface
func (b *CustomBinder) Bind(v any, c echo.Context) error {
	if err := b.d.BindQueryParams(c, v); err != nil {
		return err
	}
	return b.d.Bind(v, c)
}

// NewEcho returns a new echo instance and basic middleware is added
func NewEcho(logger *zap.Logger) *echo.Echo {
	e := echo.New()
	e.HideBanner = true // hide echo banner
	e.HidePort = true   // hide port in log

	// add prometheus middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.Use(echozap.ZapLogger(logger)) // use zap logger to replace default logger
	// add recover middleware so when panic happens, it will be recovered to centralize error handling
	e.Use(middleware.Recover())
	e.Binder = NewCustomBinder()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		logger.Error("error when handling request", zap.Error(err))
		// if the error is echo.HTTPError, it means it is a known error.
		// We can get the internal message from it.
		resp := new(types.Response)
		if he, ok := err.(*echo.HTTPError); ok {
			resp.Code = int64(he.Code)
			resp.Msg = he.Message.(string)
		} else {
			resp.Code = int64(http.StatusInternalServerError)
			resp.Msg = "failed"
		}
		// Send response
		err = c.JSON(http.StatusOK, resp)
		if err != nil {
			logger.Error("error when send response in error handler", zap.Error(err))
		}
	}
	e.Validator = api.NewCustomValidator()
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
	// add trailing slash middleware

	param.E.Pre(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		Skipper: prefixSkipper(prefix),
		// add trailing slash for all routes starting with API.Prefix of config
		// otherwise, it doesn't need trailing slash
	}))

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

// prefixSkipper returns a skipper function for middleware.
// middleware will handle the request if the request path starts with prefix
// otherwise it will skip the request
func prefixSkipper(prefix string) func(c echo.Context) bool {
	return func(c echo.Context) bool {
		if strings.HasPrefix(c.Request().URL.Path, prefix) {
			return false // don't skip
		}
		return true // otherwise skip
	}
}

// StartServer starts the HTTP server in fx.Lifecycle, so that it can be gracefully shutdown.
func StartServer(lf fx.Lifecycle, e *echo.Echo, c *config.Config) {
	server := &http.Server{Addr: c.API.Address, Handler: e}
	lf.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		}()
		return nil
	}, OnStop: func(ctx context.Context) error {
		return server.Shutdown(ctx)
	}})
}

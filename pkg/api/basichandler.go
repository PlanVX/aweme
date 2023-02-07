package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// WrapperFunc wrapper function to echo.HandlerFunc
func WrapperFunc[Req any, Resp any](biz func(context.Context, *Req) (*Resp, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(Req)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("unsupported request parameters: %w", err))
		}
		resp, err := biz(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(200, resp)
	}
}

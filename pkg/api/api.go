package api

import "github.com/labstack/echo/v4"

type Api struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

// Package api is the API layer of the application.
// It defines the API of the application.
package api

import "github.com/labstack/echo/v4"

// Api is a struct for organizing echo.HandlerFunc
type Api struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

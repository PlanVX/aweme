package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CustomValidator is the custom validator for echo
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator returns a new CustomValidator
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validates the input
func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

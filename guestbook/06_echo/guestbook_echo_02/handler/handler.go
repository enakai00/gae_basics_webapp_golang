package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
        e.HTTPErrorHandler = JSONErrorHandler
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func JSONErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	switch code {
	case 404:
		c.JSON(code, apiError{
			Status:  code,
			Message: "Error: Resource not found.",
		})
	case 500:
		c.JSON(code, apiError{
			Status:  code,
			Message: "Please contact the administrator.",
		})
	default:
		c.JSON(code, apiError{
			Status:  code,
			Message: "Unknown Error",
		})
	}
}

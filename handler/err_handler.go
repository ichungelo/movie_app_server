package handler

import (
	"go_api_echo/responses"

	"github.com/labstack/echo/v4"
)

func ErrorEndpoint(c echo.Context) (err error) {
	responses.Error404(c)
	return
}
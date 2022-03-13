package handler

import (
	"go_api_echo/repositories"
	"go_api_echo/responses"

	"github.com/labstack/echo/v4"
)

func GetAllMovies(c echo.Context) (err error)  {
	res, err := repositories.ReadAllMovies()
	if err != nil {
		responses.Error500(c)
		return
	}
	responses.StatusOK(c, res)
	return
}
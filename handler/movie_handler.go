package handler

import (
	// "go_api_echo/middlewares"
	"go_api_echo/repositories"
	"go_api_echo/responses"
	"strconv"

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

func GetMovieById(c echo.Context) (err error) {
	idString := c.Param("movieId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responses.Error404(c)
		return
	}

	if err != nil {
		responses.Error401(c, err)
		return
	}
	res, err := repositories.ReadMovieById(id)
	if err != nil {
		responses.Error404(c)
		return
	}
	responses.StatusOK(c, res)
	return
}
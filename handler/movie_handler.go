package handler

import (
	"go_api_echo/middlewares"
	"go_api_echo/repositories"
	"go_api_echo/responses"
	"log"

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
	id := c.Param("movieId")
	log.Println(id)
	payload, err := middlewares.ParseToken(c)

	if err != nil {
		responses.Error401(c, err)
		return
	}
	log.Println(payload)
	res, err := repositories.ReadMovieById(id)
	if err != nil {
		responses.Error404(c)
		return
	}
	responses.StatusOK(c, res)
	return
}
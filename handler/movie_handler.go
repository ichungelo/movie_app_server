package handler

import (
	// "go_api_echo/middlewares"
	"go_api_echo/entities"
	"go_api_echo/repositories"
	"go_api_echo/responses"
	"log"

	// "log"
	"strconv"

	"github.com/golang-jwt/jwt"
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

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtGenerateEntity)
	name := claims.UserId
	log.Println(name)

	// payload, err := middlewares.ParseToken(c)

	if err != nil {
		responses.Error401(c, err)
		return
	}
	// log.Println(payload)
	res, err := repositories.ReadMovieById(id)
	if err != nil {
		responses.Error404(c)
		return
	}
	responses.StatusOK(c, res)
	return
}
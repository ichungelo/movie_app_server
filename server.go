package main

import (
	"go_api_echo/entities"
	"go_api_echo/handler"
	"go_api_echo/middlewares"
	"go_api_echo/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}

	e.Any("*", handler.ErrorEndpoint)
	secret, err :=utils.DotEnvGenerator("SECRET_KEY")
	if err != nil {
		panic(err)
	}

	config := middleware.JWTConfig{
		Claims: &entities.JwtGenerateEntity{},
		SigningKey: []byte(secret),
	}

	accessible := e.Group("/api")
	{
		auth := accessible.Group("/auth")
		{
			auth.POST("/register", handler.RegisterUser)
			auth.POST("/login", handler.LoginUser)
		}
		accessible.GET("/movies", handler.GetAllMovies)
	}
	restricted := e.Group("/api")
	{
		restricted.Use(middleware.JWTWithConfig(config))
		restricted.GET("/movies/:movieId", handler.GetMovieById)
	}

	e.Logger.Fatal(e.Start(":5000"))
}

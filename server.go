package main

import (
	"go_api_echo/db"
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
	db.InitDB()
	db.Migrate()
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}

	e.Any("*", handler.ErrorEndpoint)
	secret, err := utils.DotEnvGenerator("SECRET_KEY")
	if err != nil {
		panic(err)
	}

	config := middleware.JWTConfig{
		Claims:     &entities.JwtGenerateEntity{},
		SigningKey: []byte(secret),
	}

	e.Static("/swaggerui", "swaggerui")

	// e.GET("test", handler.GetAllMovies)

	accessible := e.Group("/api/auth")
	{
		accessible.POST("/register", handler.RegisterUser)
		accessible.POST("/login", handler.LoginUser)
	}

	restricted := e.Group("/api/movies")
	{
		restricted.Use(middleware.JWTWithConfig(config))
		restricted.GET("/:movieId", handler.GetMovieById)
		restricted.GET("", handler.GetAllMovies)
		restricted.POST("/:movieId/review", handler.PostReview)
		restricted.PUT("/:movieId/review/:reviewId", handler.PutReview)
		restricted.DELETE("/:movieId/review/:reviewId", handler.DeleteReview)
	}

	e.Logger.Fatal(e.Start(":5000"))
}

package main

import (
	"go_api_echo/handler"
	"go_api_echo/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}

	e.Any("*", handler.ErrorEndpoint)

	e.POST("/api/register", handler.RegisterUser)
	e.POST("/api/login", handler.LoginUser)
	e.GET("/", handler.Test)
	e.Logger.Fatal(e.Start(":5000"))
}

package handler

import (
	"go_api_echo/entities"
	"go_api_echo/middlewares"
	"go_api_echo/repositories"
	"go_api_echo/responses"
	"go_api_echo/transport"
	"go_api_echo/utils"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) (err error) {
	req := new(transport.RegisterRequest)
	if err := c.Bind(req); err != nil {
		responses.Error500(c)
		return err
	}

	if err := c.Validate(req); err != nil {
		responses.Error400(c, err)
		return err
	}

	if req.Password != req.ConfirmPassword {
		responses.Error400(c, err)
		return
	}

	repositories.CreateUser(*req)

	responses.StatusOK(c, "Register Success")
	return
}

func LoginUser(c echo.Context) (err error)  {
	req := new(transport.LoginRequest)

	if err := c.Bind(req); err != nil {
		responses.Error400(c, err)
		return err
	}

	if err := c.Validate(req); err != nil {
		responses.Error400(c, err)
		return err
	}

	res := repositories.ReadUser(*req)

	passwordMatch := utils.CheckPassword(req.Password, res.Password)
	if !passwordMatch {
		responses.Error401(c, err)
		return
	}

	payload := entities.JwtGenerateEntity{
		UserId: res.UserId,
		Username: res.Username,
		Email: res.Email,
		FirstName: res.FirstName,
		LastName: res.LastName,
	}

	tokenGenrated := middlewares.GenerateToken(payload)

	responses.StatusOkLogin(c, tokenGenrated)
	return
}

func Test(c echo.Context) (err error) {
	payload, err := middlewares.ParseToken(c)
	if err != nil {
		responses.Error401(c, err)
		return
	}

	responses.StatusOK(c, payload)
	return
}
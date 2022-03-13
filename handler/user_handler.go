package handler

import (
	"errors"
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
		responses.Error400(c, errors.New("password not match"))
		return err
	}

	if err := repositories.CreateUser(*req); err != nil {
		responses.Error401(c, err)
		return err
	}

	responses.StatusOK(c, "Register Success")
	return
}

func LoginUser(c echo.Context) error {
	req := new(transport.LoginRequest)

	if err := c.Bind(req); err != nil {
		responses.Error400(c, err)
		return err
	}

	if err := c.Validate(req); err != nil {
		responses.Error400(c, err)
		return err
	}

	res, err := repositories.ReadUser(*req)
	if err != nil {
		responses.Error401(c, errors.New("invalid username or password"))
		return err
	}

	passwordMatch := utils.CheckPassword(req.Password, res.Password)
	if !passwordMatch {
		responses.Error401(c, errors.New("invalid username or password"))
		return err
	}

	payload := entities.JwtGenerateEntity{
		UserId: res.UserId,
		Username: res.Username,
		Email: res.Email,
		FirstName: res.FirstName,
		LastName: res.LastName,
	}

	tokenGenrated, err := middlewares.GenerateToken(payload)
	if err != nil {
		responses.Error500(c)
		return err
	}

	responses.StatusOkLogin(c, tokenGenrated)
	return err
}
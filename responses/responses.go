package responses

import (
	"go_api_echo/transport"
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func Error400(c echo.Context, err error) {
	res := transport.GeneralResponse{
		Success: false,
		Message: err.Error(),
	}
	 c.JSON(http.StatusBadRequest, res)
}

func Error401(c echo.Context, err error) {
	res := transport.GeneralResponse{
		Success: false,
		Message: err.Error(),
	}
	 c.JSON(http.StatusUnauthorized, res)
}

func Error404(c echo.Context) {
	res := transport.GeneralResponse{
		Success: false,
		Message: "Not Found",
	}
	 c.JSON(http.StatusNotFound, res)
}

func Error500(c echo.Context) {
	res := transport.GeneralResponse{
		Success: false,
		Message: "Internal Server Error",
	}
	 c.JSON(http.StatusInternalServerError, res)
}

//? OK
func StatusOK(c echo.Context, message interface {}) {
	res := transport.GeneralResponse{
		Success: true,
		Message: message,
	}
	c.JSON(http.StatusOK, res)
}

func StatusOkLogin(c echo.Context, token string)  {
	res :=transport.LoginResponse{
		Success: true,
		Message: "Login Success",
		Token: token,
	}
	c.JSON(http.StatusOK, res)
}



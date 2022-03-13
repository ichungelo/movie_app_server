package handler

import (
	"go_api_echo/entities"
	"go_api_echo/repositories"
	"go_api_echo/responses"
	"go_api_echo/transport"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func PostReview(c echo.Context) error  {
	movieIdString := c.Param("movieId")
	movieId, err := strconv.Atoi(movieIdString)
	if err != nil {
		responses.Error404(c)
		return err
	}

	req := new(transport.CreateReviewRequest)
	
	if err := c.Bind(req); err!= nil {
		responses.Error400(c, err)
		return err
	} 
	if err := c.Validate(req); err != nil {
		responses.Error400(c, err)
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entities.JwtGenerateEntity)
	userId := claims.UserId

	data := transport.CreateReviewResponse{
		UserId: userId,
		MovieId: movieId,
		Review: req.Review,
	}

	if err := repositories.CreateReview(data); err != nil {
		responses.Error401(c, err)
		return err
	}

	responses.StatusOK(c, "Review Added")
	return nil
}
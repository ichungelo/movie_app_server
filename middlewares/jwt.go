package middlewares

import (
	"errors"
	"fmt"
	"go_api_echo/entities"
	"go_api_echo/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GenerateToken(payload entities.JwtGenerateEntity) string {
	claims := &entities.JwtGenerateEntity{
		UserId:    payload.UserId,
		Username:  payload.Username,
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	secretKey := utils.DotEnvGenerator("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	utils.CheckError(err)

	return signedToken
}

func ParseToken(c echo.Context) (payload *entities.JwtPayloadEntity, err error) {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := strings.Split(tokenHeader, " ")

	if len(token) != 2 {
		return nil, errors.New("Unauthorized")
	}

	tokenString := token[1]
	
	if token[0] != "Bearer" {
		return nil, errors.New("invalid authorization type")
	}
	
	secretKey := []byte(utils.DotEnvGenerator("SECRET_KEY"))
	auth, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	
	
	if err != nil {
		return nil, err
	}
	if claims, ok := auth.Claims.(jwt.MapClaims); ok && auth.Valid {
		payload = &entities.JwtPayloadEntity{
			UserId:    fmt.Sprintf("%v", claims["user_id"]),
			Username:  fmt.Sprintf("%v", claims["username"]),
			Email:     fmt.Sprintf("%v", claims["email"]),
			FirstName: fmt.Sprintf("%v", claims["first_name"]),
			LastName:  fmt.Sprintf("%v", claims["last_name"]),
		}
		return payload, nil
	} else {
		utils.CheckError(err)
	}
	return
}

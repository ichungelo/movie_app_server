package utils

import (
	"go_api_echo/entities"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(payload entities.JwtGenerateEntity) (string, error) {
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

	secret, err := DotEnvGenerator("SECRET_KEY")
	if err != nil {
		return "", err
	}

	secretKey := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
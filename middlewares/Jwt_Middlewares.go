package middlewares

import (
	"Web-Golang/constants"
	"Web-Golang/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken() (string, error) {
	claims := models.JwtClaims{
		Name: "admin", // Name
		StandardClaims: jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(constants.SCREAT_JWT))
	if err != nil {
		return "", err
	}
	return token, nil
}

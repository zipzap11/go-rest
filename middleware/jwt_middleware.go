package middleware

import (
	"restful-api/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)

	return token.SignedString([]byte(constant.SECRET_JWT))
}

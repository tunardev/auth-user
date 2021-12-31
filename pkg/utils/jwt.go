package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func SignJWT(id int64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(id)),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 days
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token, err
}

func VerifyJWT(token string) (string, error) {
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return claims.Issuer, err
}
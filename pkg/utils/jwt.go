package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

const TokenExpireDuartion = time.Hour * 24

var Secret = []byte("hello_admin")

func GetToken(email string) (string, error) {
	c := MyClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuartion).Unix(),
			Issuer:    "helloadmin.cn",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

func ParseToken(t string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(t, &MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid token")
}

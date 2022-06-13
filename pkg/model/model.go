package model

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	UserId   string
	Password string
}

type JwtClaims struct {
	UserId string `json:"username"`
	jwt.StandardClaims
}

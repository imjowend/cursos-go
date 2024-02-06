package models

import "github.com/golang-jwt/jwt"

// Go utiliza composicion sobre herencia
type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

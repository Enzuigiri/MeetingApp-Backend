package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWTUserData struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}

type JwtCustomClaims struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

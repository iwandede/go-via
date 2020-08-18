package models

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type ResponseAuth struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

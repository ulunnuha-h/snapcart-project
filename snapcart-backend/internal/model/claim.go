package model

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

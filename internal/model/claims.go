package model

import "github.com/golang-jwt/jwt/v5"

// Model claims encoded to a JWT
type Claims struct {
	User string `json:"user"`
	// Provide fields like expiry time
	jwt.RegisteredClaims
}

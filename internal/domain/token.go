package domain

import "github.com/golang-jwt/jwt"

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserId int
}

package utils

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("SECRET")

type AccessToken struct {
	Value string `json:"AccessToken"`
}

type TokenResponse struct {
	Name string
	Value string
	Expires time.Time
}

func GenerateToken(tokenExpiredTime time.Time) (TokenResponse, error) {
	jwtClaims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiredTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	tokenStr, err := token.SignedString(JwtKey)
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{
		Name:    "AccessToken",
		Value:   tokenStr,
		Expires: tokenExpiredTime,
	}, nil
}

func ParseToken(tknStr string, claims *Claims) (*jwt.Token, error) {
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, Unauthorized
		}
		return nil, BadRequest
	}
	if !tkn.Valid {
		return nil, Unauthorized
	}
	return tkn, nil
}

func UserContains(slice []Creds, elem Creds) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}
package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Creds struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

var Users = []Creds{
	{
		"John@mail.io",
		"+123456",
	}, {
		"Foo@mail.com",
		"+654321",
	}}

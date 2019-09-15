package utils

import "errors"

var (
	Unauthorized = errors.New("Unauthorized")
	BadRequest = errors.New("BadRequest")
)


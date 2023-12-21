package handler

import "errors"

var (
	ErrSomethingWentWrong = errors.New("oops, Something went wrong")
	ErrInvalidInput       = errors.New("invalid input")
)

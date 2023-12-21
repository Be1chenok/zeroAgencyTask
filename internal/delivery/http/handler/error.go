package handler

import "errors"

var (
	ErrSomethingWentWrong = errors.New("oops, Something went wrong")
	ErrInvalidInput       = errors.New("invalid input")
	ErrEmptyHeader        = errors.New("empty auth header")
	ErrInvalidAuthHeader  = errors.New("invalid auth header")
	ErrEmptyToken         = errors.New("empty token")
)

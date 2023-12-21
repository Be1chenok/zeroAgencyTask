package domain

import "errors"

var (
	ErrNothingUpdated = errors.New("nothing updated")
	ErrNothingFound   = errors.New("nothing found")
)

package handler

import "github.com/Be1chenok/zeroAgencyTask/internal/domain"

type listNews struct {
	Success bool           `json:"Success"`
	News    *[]domain.News `json:"News"`
}

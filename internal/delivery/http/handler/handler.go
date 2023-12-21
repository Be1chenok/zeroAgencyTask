package handler

import (
	"github.com/Be1chenok/zeroAgencyTask/internal/config"
	appService "github.com/Be1chenok/zeroAgencyTask/internal/service"
)

type Handler struct {
	service *appService.Service
	conf    *config.Config
}

func New(conf *config.Config, service *appService.Service) *Handler {
	return &Handler{
		service: service,
		conf:    conf,
	}
}

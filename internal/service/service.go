package service

import (
	appLogger "github.com/Be1chenok/zeroAgencyTask/internal/logger"
	"github.com/Be1chenok/zeroAgencyTask/internal/repository"
)

type Service struct {
	News
}

func New(repo *repository.Repository, logger appLogger.Logger) *Service {
	return &Service{
		News: NewNews(repo.News, logger),
	}
}

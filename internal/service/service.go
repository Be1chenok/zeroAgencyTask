package service

import (
	"github.com/Be1chenok/zeroAgencyTask/internal/config"
	appLogger "github.com/Be1chenok/zeroAgencyTask/internal/logger"
	"github.com/Be1chenok/zeroAgencyTask/internal/repository"
	appHasher "github.com/Be1chenok/zeroAgencyTask/pkg/hasher"
)

type Service struct {
	News
	Authentication
}

func New(repo *repository.Repository, logger appLogger.Logger, hasher *appHasher.SHA256Hasher, conf *config.Config) *Service {
	return &Service{
		News:           NewNews(repo.News, logger),
		Authentication: NewAuthentication(repo.User, hasher, conf),
	}
}

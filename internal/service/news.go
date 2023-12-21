package service

import (
	"context"
	"fmt"

	"github.com/Be1chenok/zeroAgencyTask/internal/domain"
	appLogger "github.com/Be1chenok/zeroAgencyTask/internal/logger"
	"github.com/Be1chenok/zeroAgencyTask/internal/repository/postgres"
	"go.uber.org/zap"
)

type News interface {
	FindNews(ctx context.Context, searchParams domain.NewsSearchParams) (*[]domain.News, error)
	UpdateNewsById(ctx context.Context, news *domain.News) error
}

type news struct {
	postgresNews postgres.News
	logger       appLogger.Logger
}

func NewNews(postgresNews postgres.News, logger appLogger.Logger) News {
	return &news{
		postgresNews: postgresNews,
		logger:       logger.With(zap.String("component", "service-news")),
	}
}

func (n news) FindNews(ctx context.Context, searchParams domain.NewsSearchParams) (*[]domain.News, error) {
	newsList, err := n.postgresNews.GetListNews(ctx, searchParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get news list: %w", err)
	}

	return newsList, err
}

func (n news) UpdateNewsById(ctx context.Context, news *domain.News) error {
	if err := n.postgresNews.EditNewsById(ctx, news); err != nil {
		return fmt.Errorf("failed to update news by id: %w", err)
	}

	return nil
}

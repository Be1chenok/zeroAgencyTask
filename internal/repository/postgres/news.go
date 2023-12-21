package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Be1chenok/zeroAgencyTask/internal/domain"
	appLogger "github.com/Be1chenok/zeroAgencyTask/internal/logger"
	"go.uber.org/zap"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type News interface {
	GetListNews(ctx context.Context, searchParams domain.NewsSearchParams) (*[]domain.News, error)
	EditNewsById(ctx context.Context, news *domain.News) error
}

type news struct {
	logger appLogger.Logger
	db     *reform.DB
}

func NewNewsRepo(db *sql.DB, logger appLogger.Logger) News {
	logger = logger.With(zap.String("component", "postgres-news"))
	return &news{
		logger: logger,
		db:     reform.NewDB(db, postgresql.Dialect, reform.NewPrintfLogger(logger.Infof)),
	}
}

func wrapCommitError(err, e error) error {
	return fmt.Errorf("commit tx: %w:%w", err, e)
}

func wrapBeginError(err error) error {
	return fmt.Errorf("begin tx: %w", err)
}

func wrapRollbackError(err, e error) error {
	return fmt.Errorf("rollback tx: %w:%w", err, e)
}

func (n news) EditNewsById(ctx context.Context, news *domain.News) error {
	tx, err := n.db.BeginTx(ctx, nil)
	if err != nil {
		return wrapBeginError(err)
	}
	defer func() {
		if err != nil {
			if e := tx.Rollback(); e != nil {
				err = wrapRollbackError(err, e)

				return
			}

			return
		}

		if e := tx.Commit(); e != nil {
			err = wrapCommitError(err, e)
		}
	}()

	result, err := tx.Exec(`
		UPDATE news
		SET title = $1,
		content = $2
		WHERE id = $3`,
		news.Title,
		news.Content,
		news.Id)
	if err != nil {
		return fmt.Errorf("failed to execute UPDATE query: %w", err)
	}

	updatedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get number of updated rows: %w", err)
	}
	if updatedRows == 0 {
		return domain.ErrNothingUpdated
	}

	if news.Categories != nil {
		if _, err := tx.DeleteFrom(domain.NewsCategoriesTable, "WHERE news_id=$1", news.Id); err != nil {
			return fmt.Errorf("remove existing categories for news_id '%d': %w", news.Id, err)
		}

		for _, element := range news.Categories {
			category := &domain.NewsCategories{
				NewsId:     news.Id,
				CategoryId: element,
			}

			if err := tx.Save(category); err != nil {
				return fmt.Errorf("failed to save news category: %w", err)
			}
		}
	}

	return nil
}

func (n news) GetListNews(ctx context.Context, searchParams domain.NewsSearchParams) (*[]domain.News, error) {
	tx, err := n.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, wrapBeginError(err)
	}
	defer func() {
		if err != nil {
			if e := tx.Rollback(); e != nil {
				err = wrapRollbackError(err, e)

				return
			}

			return
		}

		if e := tx.Commit(); e != nil {
			err = wrapCommitError(err, e)
		}
	}()

	newsList, err := tx.SelectAllFrom(domain.NewsTable, `LIMIT $1 OFFSET $2`, searchParams.Limit, searchParams.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute select query from news table: %w", err)
	}
	if len(newsList) == 0 {
		return nil, domain.ErrNothingFound
	}

	result := make([]domain.News, 0, len(newsList))
	for _, element := range newsList {
		news := *element.(*domain.News)

		categories, err := tx.SelectAllFrom(domain.NewsCategoriesTable, "WHERE news_id = $1", news.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to execute select query from news_categories table: %w", err)
		}

		for _, category := range categories {
			news.Categories = append(news.Categories, category.(*domain.NewsCategories).CategoryId)
		}

		result = append(result, news)
	}

	return &result, nil
}

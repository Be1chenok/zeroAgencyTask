package repository

import (
	"database/sql"

	appLogger "github.com/Be1chenok/zeroAgencyTask/internal/logger"
	"github.com/Be1chenok/zeroAgencyTask/internal/repository/postgres"
)

type Repository struct {
	News postgres.News
	User postgres.User
}

func New(logger appLogger.Logger, db *sql.DB) *Repository {
	return &Repository{
		News: postgres.NewNewsRepo(db, logger),
		User: postgres.NewUserRepo(db, logger),
	}
}

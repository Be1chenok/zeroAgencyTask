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

type User interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)

	SetTokens(ctx context.Context, userId int, tokens *domain.Tokens) error

	GetUserId(ctx context.Context, username, passwordHash string) (int, error)
	GetUserIdByAccessToken(ctx context.Context, accessToken string) (int, error)
	GetUserIdByRefreshToken(ctx context.Context, refreshToken string) (int, error)

	DeleteUserIdByAccessToken(ctx context.Context, accessToken string) error
	DeleteUserIdByRefreshToken(ctx context.Context, refreshToken string) error
	DeleteAllTokensByUserId(ctx context.Context, userId int) error
}

type user struct {
	logger appLogger.Logger
	db     *reform.DB
}

func NewUserRepo(db *sql.DB, logger appLogger.Logger) User {
	logger = logger.With(zap.String("component", "postgres-user"))
	return &user{
		logger: logger,
		db:     reform.NewDB(db, postgresql.Dialect, reform.NewPrintfLogger(logger.Infof)),
	}
}

func (u user) CreateUser(ctx context.Context, user domain.User) (int, error) {
	var userId int
	if err := u.db.QueryRowContext(
		ctx,
		`INSERT INTO users (email, username, password)
		values ($1, $2, $3) RETURNING id`,
		user.Email,
		user.Username,
		user.Password,
	).Scan(&userId); err != nil {
		return 0, fmt.Errorf("failed to scan row: %w", err)
	}

	return userId, nil
}

func (u user) GetUserId(ctx context.Context, username string, passwordHash string) (int, error) {
	var userId int

	if err := u.db.QueryRowContext(
		ctx,
		`SELECT id FROM users WHERE username=$1 AND password=$2`,
		username, passwordHash,
	).Scan(&userId); err != nil {
		return 0, domain.ErrNothingFound
	}

	return userId, nil
}

func (u user) GetUserIdByAccessToken(ctx context.Context, accessToken string) (int, error) {
	var userId int

	if err := u.db.QueryRowContext(
		ctx,
		`SELECT user_id FROM tokens WHERE access_token=$1`,
		accessToken,
	).Scan(&userId); err != nil {
		return 0, domain.ErrNothingFound
	}

	return userId, nil
}

func (u user) GetUserIdByRefreshToken(ctx context.Context, refreshToken string) (int, error) {
	var userId int

	if err := u.db.QueryRowContext(
		ctx,
		`SELECT user_id FROM tokens WHERE refresh_token=$1`,
		refreshToken,
	).Scan(&userId); err != nil {
		return 0, domain.ErrNothingFound
	}

	return userId, nil
}

func (u user) SetTokens(ctx context.Context, userId int, tokens *domain.Tokens) error {
	if _, err := u.db.ExecContext(
		ctx,
		`INSERT INTO tokens (user_id, access_token, refresh_token) values ($1,$2,$3)`,
		userId,
		tokens.AccessToken,
		tokens.RefreshToken,
	); err != nil {
		return fmt.Errorf("failed to execute insert query id tokens table: %w", err)
	}

	return nil
}

func (u user) DeleteUserIdByAccessToken(ctx context.Context, accessToken string) error {
	if _, err := u.db.ExecContext(
		ctx,
		`DELETE FROM tokens WHERE access_token=$1`,
		accessToken,
	); err != nil {
		return fmt.Errorf("failed to execute delete query in tokens table: %w", err)
	}

	return nil
}

func (u user) DeleteUserIdByRefreshToken(ctx context.Context, refreshToken string) error {
	if _, err := u.db.ExecContext(
		ctx,
		`DELETE FROM tokens WHERE refresh_token=$1`,
		refreshToken,
	); err != nil {
		return fmt.Errorf("failed to execute delete query in tokens table: %w", err)
	}

	return nil
}

func (u user) DeleteAllTokensByUserId(ctx context.Context, userId int) error {
	if _, err := u.db.ExecContext(
		ctx,
		`DELETE FROM tokens WHERE user_id=$1`,
		userId,
	); err != nil {
		return fmt.Errorf("failed to execute delete query in tokens table: %w", err)
	}

	return nil
}

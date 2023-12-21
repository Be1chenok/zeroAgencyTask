package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Be1chenok/zeroAgencyTask/internal/config"
	"github.com/Be1chenok/zeroAgencyTask/internal/domain"
	"github.com/Be1chenok/zeroAgencyTask/internal/repository/postgres"
	appHasher "github.com/Be1chenok/zeroAgencyTask/pkg/hasher"
	"github.com/golang-jwt/jwt"
)

type Authentication interface {
	SignUp(ctx context.Context, input domain.User) (int, error)
	SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error)

	RefreshTokens(ctx context.Context, refreshToken string) (*domain.Tokens, error)
	ParseToken(ctx context.Context, accessToken string) (int, error)

	SignOut(ctx context.Context, accessToken string) error
	FullSignOut(ctx context.Context, accessToken string) error
}

type authentication struct {
	postgresUser postgres.User
	hasher       *appHasher.SHA256Hasher
	conf         *config.Config
}

func NewAuthentication(postgresUser postgres.User, hasher *appHasher.SHA256Hasher, conf *config.Config) Authentication {
	return &authentication{
		postgresUser: postgresUser,
		hasher:       hasher,
		conf:         conf,
	}
}

func (a authentication) SignUp(ctx context.Context, input domain.User) (int, error) {
	passwordHash, err := a.hasher.Hash(input.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	user := domain.User{
		Email:    input.Email,
		Username: input.Username,
		Password: passwordHash,
	}

	userId, err := a.postgresUser.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return userId, nil
}

func (a authentication) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	passwordHash, err := a.hasher.Hash(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	userId, err := a.postgresUser.GetUserId(ctx, input.Username, passwordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id: %w", err)
	}

	tokens, err := a.createSession(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return tokens, nil
}

func (a authentication) SignOut(ctx context.Context, accessToken string) error {
	if err := a.postgresUser.DeleteUserIdByAccessToken(ctx, accessToken); err != nil {
		return fmt.Errorf("failed to delete user id by access token: %w", err)
	}

	return nil
}

func (a authentication) FullSignOut(ctx context.Context, accessToken string) error {
	userId, err := a.postgresUser.GetUserIdByAccessToken(ctx, accessToken)
	if err != nil {
		return fmt.Errorf("failed to get user id by access token: %w", err)
	}

	if err := a.postgresUser.DeleteAllTokensByUserId(ctx, userId); err != nil {
		return fmt.Errorf("failed to delete all tokens by user id: %w", err)
	}

	return nil
}

func (a authentication) RefreshTokens(ctx context.Context, refreshToken string) (*domain.Tokens, error) {
	userId, err := a.postgresUser.GetUserIdByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id by refresh token: %w", err)
	}
	if err := a.postgresUser.DeleteUserIdByRefreshToken(ctx, refreshToken); err != nil {
		return nil, fmt.Errorf("failed to delete user id by refresh token: %w", err)
	}

	tokens, err := a.createSession(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return tokens, nil
}

func (a authentication) ParseToken(ctx context.Context, token string) (int, error) {
	userId, err := a.postgresUser.GetUserIdByAccessToken(ctx, token)
	if err != nil {
		userId, err = a.postgresUser.GetUserIdByRefreshToken(ctx, token)
		if err != nil {
			return 0, fmt.Errorf("failed to get user id by token: %w", err)
		}
	}

	parsedToken, err := jwt.ParseWithClaims(token, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(a.conf.Tokens.SigningKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to parsing token: %w", err)
	}

	claims, ok := parsedToken.Claims.(*domain.AccessTokenClaims)
	if !ok {
		return 0, fmt.Errorf("token claims are not of type *tokenClaims")
	}

	if claims.UserId != userId {
		return 0, fmt.Errorf("token does not match the stored user id")
	}

	return claims.UserId, nil
}

func (a authentication) createSession(ctx context.Context, userId int) (*domain.Tokens, error) {
	accessToken, err := a.createToken(userId, a.conf.Tokens.AccessTTL)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	refreshToken, err := a.createToken(userId, a.conf.Tokens.RefreshTTL)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}

	tokens := domain.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := a.postgresUser.SetTokens(ctx, userId, &tokens); err != nil {
		return nil, fmt.Errorf("failed to set access and refresh tokens: %w", err)
	}

	return &tokens, nil
}

func (a authentication) createToken(userId int, tokenTTL time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})

	signedToken, err := token.SignedString([]byte(a.conf.Tokens.SigningKey))
	if err != nil {
		return "", fmt.Errorf("failed to signing token: %w", err)
	}

	return signedToken, nil
}

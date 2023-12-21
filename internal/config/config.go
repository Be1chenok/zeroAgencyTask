package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server       ServerConfig
	Postgres     PostgresConfig
	UserPassword UserPasswordConfig
	Tokens       TokensConfig
}

type ServerConfig struct {
	Host           string
	Port           int
	RequestTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type UserPasswordConfig struct {
	Salt string
}

type TokensConfig struct {
	SigningKey string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func Init() (*Config, error) {
	viper.SetConfigFile("../../.env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		ServerConfig{
			Host:           viper.GetString("SERVER_HOST"),
			Port:           viper.GetInt("SERVER_PORT"),
			RequestTimeout: viper.GetDuration("SERVER_REQUEST_TIMEOUT") * time.Second,
			ReadTimeout:    viper.GetDuration("SERVER_READ_TIMEOUT") * time.Second,
			WriteTimeout:   viper.GetDuration("SERVER_WRITE_TIMEOUT") * time.Second,
		},
		PostgresConfig{
			Host:     viper.GetString("PG_HOST"),
			Port:     viper.GetInt("PG_PORT"),
			Username: viper.GetString("PG_USER"),
			Password: viper.GetString("PG_PASS"),
			DBName:   viper.GetString("PG_BASE"),
			SSLMode:  viper.GetString("PG_SSL_MODE"),
		},
		UserPasswordConfig{
			Salt: viper.GetString("USER_PASSWORD_SALT"),
		},
		TokensConfig{
			SigningKey: viper.GetString("TOKENS_SIGNING_KEY"),
			AccessTTL:  viper.GetDuration("ACCESS_TOKEN_TTL") * time.Second,
			RefreshTTL: viper.GetDuration("REFRESH_TOKEN_TTL") * time.Second,
		},
	}, nil
}

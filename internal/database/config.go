package database

import (
	"errors"
	"fmt"
)

const (
	emptyHost       = `empty host`
	emptyPort       = `empty port`
	emptyUserName   = `empty user name`
	emptyPassword   = `empty password`
	emptyDBName     = `empty db name`
	emptyTestDBName = `empty test db name`
	emptySSLMode    = `empty SSL mode`
)

const (
	EnvHost       = "CAR_DB_HOST"
	EnvPort       = "CAR_DB_PORT"
	EnvUserName   = "CAR_DB_USER_NAME"
	EnvPassword   = "CAR_DB_PASSWORD"
	EnvDBName     = "CAR_DB_NAME"
	EnvTestDBName = "CAR_TEST_DB_NAME"
	EnvSSLMode    = "CAR_DB_SSL_MODE"
)

var (
	ErrEmptyHost       = errors.New(emptyHost)
	ErrEmptyPort       = errors.New(emptyPort)
	ErrEmptyUserName   = errors.New(emptyUserName)
	ErrEmptyPassword   = errors.New(emptyPassword)
	ErrEmptyDBName     = errors.New(emptyDBName)
	ErrEmptyTestDBName = errors.New(emptyTestDBName)
	ErrEmptySSLMode    = errors.New(emptySSLMode)
)

type Config struct {
	Host       string
	Port       string
	UserName   string
	Password   string
	DBName     string
	TestDBName string
	SSLMode    string
}

// Validate checks if the config is ok
func (cfg *Config) Validate() error {

	if len(cfg.Host) == 0 {
		return ErrEmptyHost
	}

	if len(cfg.Port) == 0 {
		return ErrEmptyPort
	}

	if len(cfg.UserName) == 0 {
		return ErrEmptyUserName
	}

	if len(cfg.Password) == 0 {
		return ErrEmptyPassword
	}

	if len(cfg.DBName) == 0 {
		return ErrEmptyDBName
	}

	if len(cfg.TestDBName) == 0 {
		return ErrEmptyTestDBName
	}

	if len(cfg.SSLMode) == 0 {
		return ErrEmptySSLMode
	}

	return nil
}

// String returns prepared connection DB URL
func (cfg *Config) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
}

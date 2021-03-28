package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/lks-go/car/internal/server"

	"github.com/lks-go/car/internal/database"
)

func New() *config {
	return &config{}
}

type config struct {
	Server   *server.Config
	Database *database.Config
}

func (cfg *config) Init() error {

	serverPort := os.Getenv(server.EnvPort)
	if len(serverPort) == 0 {
		serverPort = server.DefaultPort
	}
	cfg.Server = &server.Config{
		Port: serverPort,
	}

	cfg.Database = &database.Config{
		Host:       os.Getenv(database.EnvHost),
		Port:       os.Getenv(database.EnvPort),
		UserName:   os.Getenv(database.EnvUserName),
		Password:   os.Getenv(database.EnvPassword),
		DBName:     os.Getenv(database.EnvDBName),
		TestDBName: os.Getenv(database.EnvTestDBName),
		SSLMode:    os.Getenv(database.EnvSSLMode),
	}

	if err := cfg.validate(); err != nil {
		return err
	}

	return nil
}

func (cfg *config) validate() error {

	if err := cfg.Database.Validate(); err != nil {
		return errors.New(fmt.Sprintf("database config: %s", err.Error()))
	}

	return nil
}

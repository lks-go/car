package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Start(cfg *Config, h http.Handler) error {
	addr := cfg.Addr()

	log.Printf("listen and serve %s", addr)

	srv := &http.Server{
		Addr:    cfg.Addr(),
		Handler: h,
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit

		srv.Shutdown(context.Background())

		log.Println("shutting down...")
	}()

	return srv.ListenAndServe()
}

const (
	EnvPort = "CAR_SERVER_PORT"

	DefaultPort = "8080"
)

type Config struct {
	Port string
}

func (cfg *Config) Addr() string {
	return ":" + cfg.Port
}

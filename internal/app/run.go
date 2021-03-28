package app

import (
	"fmt"
	"log"

	"github.com/lks-go/car/internal/router"

	"github.com/lks-go/car/internal/delivery/handler"

	"github.com/lks-go/car/internal/server"

	"github.com/lks-go/car/internal/database"

	"github.com/lks-go/car/internal/repository"

	"github.com/lks-go/car/internal/config"
)

// Run inits configs, database, repository, handlers and starts http server
func Run() {
	log.Println("initializing configs...")
	cfg := config.New()
	if err := cfg.Init(); err != nil {
		log.Fatal(fmt.Errorf("can't initialize config: %s", err))
	}

	log.Println("connecting to database...")
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal(fmt.Errorf("can't connect to db: %s", err))
	}
	defer db.Close()

	repo := repository.New(db)

	r := router.InitRoutes(&router.Handlers{
		Car: handler.NewCarHandler(repo.Car),
	})

	if err := server.Start(cfg.Server, r); err != nil {
		log.Println(err)
	}
}

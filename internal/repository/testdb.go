package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/lks-go/car/internal/database"

	"github.com/lks-go/car/internal/config"
)

func TestDB(t *testing.T) (*sql.DB, func(...string)) {
	t.Helper()

	cfg := config.New()
	if err := cfg.Init(); err != nil {
		log.Fatal(fmt.Errorf("can't initialize config: %s", err))
	}

	cfg.Database.DBName = os.Getenv(database.EnvTestDBName)

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal(fmt.Errorf("can't connect to db: %s", err))
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}

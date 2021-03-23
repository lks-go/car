package database

import (
	"context"
	"database/sql"
	"fmt"
)

// InitSchema creates needed tables only to simplify the project demonstration
func InitSchema(db *sql.DB) (err error) {

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer func() {
		if e := tx.Rollback(); e == nil {
			err = fmt.Errorf("rolled back becouse error occured: %s", err)
		}
	}()

	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS car
		(
			id    SERIAL PRIMARY KEY,
			brand VARCHAR NOT NULL,
			model VARCHAR NOT NULL,
			price INTEGER NOT NULL,
			status VARCHAR,
			mileage INTEGER
		);`)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

package repository

import (
	"database/sql"

	"github.com/lks-go/car/internal/repository/car"
)

func New(db *sql.DB) *Repository {
	return &Repository{
		Car: car.New(db),
	}
}

type Repository struct {
	Car *car.Repository
}

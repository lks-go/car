package car

import (
	"database/sql"
	"fmt"

	"github.com/lks-go/car/internal/database"

	"github.com/lks-go/car/internal/domain"
)

// New returns the repository pointer
func New(db *sql.DB) *Repository {
	return &Repository{db}
}

// Repository contains methods for working with car data
type Repository struct {
	db *sql.DB
}

func (r *Repository) Create(c *domain.Car) (*domain.Car, error) {
	newCar := &domain.Car{}
	q := fmt.Sprintf(
		`INSERT INTO %s (brand, model, price, status, mileage) VALUES ($1, $2, $3, $4, $5) RETURNING id, brand, model, price, status, mileage`,
		database.TableCar,
	)
	err := r.db.QueryRow(
		q, c.Brand, c.Model, c.Price, c.Status, c.Mileage,
	).Scan(&newCar.ID, &newCar.Brand, &newCar.Model, &newCar.Price, &newCar.Status, &newCar.Mileage)
	if err != nil {
		return nil, err
	}

	return newCar, nil
}

func (r *Repository) GetByID(ID uint) (*domain.Car, error) {
	car := &domain.Car{}

	q := fmt.Sprintf(
		`SELECT id, brand, model, price, status, mileage FROM %s WHERE id = $1`,
		database.TableCar,
	)

	err := r.db.QueryRow(q, ID).Scan(&car.ID, &car.Brand, &car.Model, &car.Price, &car.Status, &car.Mileage)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return car, nil
}

func (r *Repository) Update(c *domain.Car) (*domain.Car, error) {
	car := &domain.Car{}

	q := fmt.Sprintf(
		`UPDATE %s SET
			price = $2,
			status = $3,
			mileage = $4
		WHERE id = $1
		RETURNING id, brand, model, price, status, mileage`,
		database.TableCar,
	)

	err := r.db.QueryRow(q, c.ID, c.Price, c.Status, c.Mileage).Scan(&car.ID, &car.Brand, &car.Model, &car.Price, &car.Status, &car.Mileage)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (r *Repository) Delete(ID uint) error {
	q := fmt.Sprintf(
		`DELETE FROM %s WHERE id = $1`,
		database.TableCar,
	)

	_, err := r.db.Exec(q, ID)
	if err != nil {
		return err
	}

	return nil
}

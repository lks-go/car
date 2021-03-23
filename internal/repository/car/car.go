package car

import (
	"database/sql"

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

	err := r.db.QueryRow(`
		INSERT INTO car (brand, model, price, status, mileage)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, brand, model, price, status, mileage;`,
		c.Brand, c.Model, c.Price, c.Status, c.Mileage,
	).Scan(&newCar.ID, &newCar.Brand, &newCar.Model, &newCar.Price, &newCar.Status, &newCar.Mileage)
	if err != nil {
		return nil, err
	}

	return newCar, nil
}

func (r *Repository) GetByID(ID uint) (*domain.Car, error) {
	car := &domain.Car{}

	err := r.db.QueryRow(`
		SELECT id, brand, model, price, status, mileage FROM car WHERE id = $1;`,
		ID,
	).Scan(&car.ID, &car.Brand, &car.Model, &car.Price, &car.Status, &car.Mileage)
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

	err := r.db.QueryRow(`
		UPDATE car SET
			price = $2,
			status = $3,
			mileage = $4
		WHERE id = $1
		RETURNING id, brand, model, price, status, mileage;
	`, c.ID, c.Price, c.Status, c.Mileage).Scan(&car.ID, &car.Brand, &car.Model, &car.Price, &car.Status, &car.Mileage)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (r *Repository) Delete(ID uint) error {
	_, err := r.db.Exec(`DELETE FROM car WHERE id = $1;`, ID)
	if err != nil {
		return err
	}

	return nil
}

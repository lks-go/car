package car_test

import (
	"testing"

	"github.com/lks-go/car/internal/repository/car"
	"github.com/stretchr/testify/assert"

	"github.com/lks-go/car/internal/domain"

	"github.com/lks-go/car/internal/database"

	"github.com/lks-go/car/internal/repository"
)

var newCar = &domain.Car{
	Brand:   "Hyundai",
	Model:   "SOLARIS",
	Price:   820000,
	Status:  domain.InStock,
	Mileage: 0,
}

func TestCreateCar(t *testing.T) {
	db, truncate := repository.TestDB(t)
	defer db.Close()
	defer truncate(database.TableCar)

	repo := car.New(db)
	createdCar, err := repo.Create(newCar)

	assert.Nil(t, err)
	assert.NotNil(t, createdCar)
}

func TestGetCarByID(t *testing.T) {
	db, truncate := repository.TestDB(t)
	defer db.Close()
	defer truncate(database.TableCar)

	repo := car.New(db)
	createdCar, err := repo.Create(newCar)
	assert.Nil(t, err)

	gotCar, err := repo.GetByID(createdCar.ID)
	assert.Nil(t, err)
	assert.NotNil(t, gotCar)

	gotCar, err = repo.GetByID(createdCar.ID + 1)
	assert.Nil(t, err)
	assert.Nil(t, gotCar)
}

func TestUpdateCar(t *testing.T) {
	db, truncate := repository.TestDB(t)
	defer db.Close()
	defer truncate(database.TableCar)

	repo := car.New(db)
	createdCar, err := repo.Create(newCar)
	assert.Nil(t, err)

	createdCar.Status = domain.InTransit

	gotCar, err := repo.Update(createdCar)
	assert.Nil(t, err)
	assert.NotNil(t, gotCar)
	assert.Equal(t, *createdCar, *gotCar)
}

func TestDeleteCar(t *testing.T) {
	db, truncate := repository.TestDB(t)
	defer db.Close()
	defer truncate(database.TableCar)

	repo := car.New(db)
	createdCar, err := repo.Create(newCar)
	assert.Nil(t, err)

	err = repo.Delete(createdCar.ID)
	assert.Nil(t, err)

	gotCar, err := repo.GetByID(createdCar.ID)
	assert.Nil(t, err)
	assert.Nil(t, gotCar)
}

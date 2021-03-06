package handler

import (
	"errors"
	"fmt"

	"github.com/lks-go/car/internal/domain"
)

func NewCarMock() *CarMock {

	cm := CarMock{
		db: make(map[uint]domain.Car),
	}

	cm.db[1] = domain.Car{
		ID:      1,
		Brand:   "Mitsubishi",
		Model:   "Carisma",
		Price:   200000,
		Status:  domain.Sold,
		Mileage: 190000,
	}

	return &cm
}

type CarMock struct {
	db map[uint]domain.Car
}

func (cm *CarMock) GetByID(ID uint) (*domain.Car, error) {
	car, ok := cm.db[ID]
	if !ok {
		return nil, nil
	}

	return &car, nil
}

func (cm *CarMock) Create(c *domain.Car) (*domain.Car, error) {
	newCarID := uint(len(cm.db) + 1)

	c.ID = newCarID

	cm.db[newCarID] = *c

	newCar, ok := cm.db[newCarID]
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &newCar, nil
}

func (cm *CarMock) Update(c *domain.Car) (*domain.Car, error) {

	if _, ok := cm.db[c.ID]; !ok {
		return nil, nil
	}

	cm.db[c.ID] = *c

	updatedCar := cm.db[c.ID]

	return &updatedCar, nil
}

func (cm *CarMock) Delete(ID uint) error {
	if _, ok := cm.db[ID]; !ok {
		return errors.New(fmt.Sprintf("record with ID %d not found", ID))
	}

	delete(cm.db, ID)

	return nil
}

package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/lks-go/car/internal/domain"
)

type Car interface {
	GetByID(ID uint) (*domain.Car, error)
	Create(c *domain.Car) (*domain.Car, error)
	Update(c *domain.Car) (*domain.Car, error)
	Delete(ID uint) error
}

func NewCarHandler(repo Car) *CarHandler {
	return &CarHandler{
		Car: repo,
	}
}

type CarHandler struct {
	Car Car
}

func (h *CarHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	car, err := h.Car.GetByID(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	if car == nil {
		responseError(w, http.StatusNotFound, nil)
		return
	}

	b, err := json.Marshal(car)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	responseOk(w, b)
}

func (h *CarHandler) Create(w http.ResponseWriter, r *http.Request) {
	car := domain.Car{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(body, &car); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	newCar, err := h.Car.Create(&car)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	responseCreated(w, fmt.Sprintf("/api/v1/car/%d", newCar.ID))
}

func (h *CarHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	car, err := h.Car.GetByID(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	if car == nil {
		responseError(w, http.StatusNotFound, nil)
		return
	}

	car = &domain.Car{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(body, &car); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	car.ID = uint(id)
	car, err = h.Car.Update(car)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	b, err := json.Marshal(car)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	responseOk(w, b)
}

func (h *CarHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	car, err := h.Car.GetByID(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	if car == nil {
		responseError(w, http.StatusNotFound, nil)
		return
	}

	err = h.Car.Delete(uint(id))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	responseOk(w, nil)
}

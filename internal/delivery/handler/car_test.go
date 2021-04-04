// +build handlers

package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lks-go/car/internal/domain"

	"github.com/gorilla/mux"

	"github.com/lks-go/car/internal/delivery/handler"
)

func TestHandlerGetByID(t *testing.T) {

	testCases := []struct {
		ID         uint
		shouldPass bool
	}{
		{1, true},
		{0, true},
		{2, false},
	}

	h := handler.NewCarHandler(handler.NewCarMock())

	for _, tc := range testCases {
		path := fmt.Sprintf("/api/v1/car/%d", tc.ID)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", path, nil)
		r := mux.NewRouter()

		r.HandleFunc("/api/v1/car/{id:[0-9]+}", h.GetByID)
		r.ServeHTTP(rec, req)

		if rec.Code == http.StatusOK && !tc.shouldPass {
			t.Errorf("handler should have failed on routeVariable %d: got %v",
				tc.ID, rec.Code)
		}
	}
}

func TestHandlerCreate(t *testing.T) {
	h := handler.NewCarHandler(handler.NewCarMock())

	car := domain.Car{
		Brand:   "BMW",
		Model:   "X5",
		Price:   3000000,
		Status:  domain.InStock,
		Mileage: 55000,
	}

	b, err := json.Marshal(car)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/car/", bytes.NewReader(b))
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/car/", h.Create)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("handler should have return http code %d: got %d", http.StatusCreated, rec.Code)
		return
	}

	newCarLocation := rec.Header().Get("Location")
	rec = httptest.NewRecorder()
	req = httptest.NewRequest("GET", newCarLocation, nil)
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/car/{id:[0-9]+}", h.GetByID)
	r.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("handler should have return http code %d: got %d", http.StatusOK, rec.Code)
		return
	}
}

func TestHandlerUpdate(t *testing.T) {
	mock := handler.NewCarMock()
	h := handler.NewCarHandler(mock)

	car, err := mock.GetByID(1)
	if err != nil {
		t.Errorf("didn't expect an error: got %v", err)
		return
	}

	// changing car data
	car.Mileage = 210000
	car.Price = 180000

	b, err := json.Marshal(car)
	if err != nil {
		t.Errorf("didn't expect an error: got %v", err)
		return
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/api/v1/car/1", bytes.NewReader(b))
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/car/{id:[0-9]+}", h.Update)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("handler should have return http code %d: got %d", http.StatusOK, rec.Code)
		return
	}

	resCar := domain.Car{}
	if err := json.NewDecoder(rec.Body).Decode(&resCar); err != nil {
		t.Errorf("didn't expect an error: got %v", err)
		return
	}

	if resCar.ID != car.ID || resCar.Price != car.Price || resCar.Mileage != car.Mileage {
		t.Errorf("expected ID: %v, Price: %v, Mileage: %v\n got ID: %v, Price: %v, Mileage: %v",
			car.ID, car.Price, car.Mileage, resCar.ID, resCar.Price, resCar.Mileage)
	}

}

func TestHandlerDelete(t *testing.T) {
	testCases := []struct {
		ID                 uint
		ExpectedStatusCode int
		ShouldPass         bool
	}{
		{ID: 1, ExpectedStatusCode: http.StatusOK, ShouldPass: true},
		{ID: 100, ExpectedStatusCode: http.StatusNotFound, ShouldPass: true},
		{ID: 1, ExpectedStatusCode: http.StatusOK, ShouldPass: false},
	}

	h := handler.NewCarHandler(handler.NewCarMock())

	for _, tc := range testCases {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/car/%d", tc.ID), nil)
		r := mux.NewRouter()
		r.HandleFunc("/api/v1/car/{id:[0-9]+}", h.Delete)
		r.ServeHTTP(rec, req)

		if tc.ShouldPass && rec.Code != tc.ExpectedStatusCode {
			t.Errorf("handler should have return http code %d: got %d", tc.ExpectedStatusCode, rec.Code)
		}
	}

}

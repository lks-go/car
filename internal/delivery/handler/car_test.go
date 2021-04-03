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

}

func TestHandlerDelete(t *testing.T) {

}

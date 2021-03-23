package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes initialize http paths for handlers
func InitRoutes(h *Handlers) http.Handler {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/api/v1").Subrouter()

	v1.Path("/car/").HandlerFunc(h.Car.Create).Methods("POST")
	v1.Path("/car/{id:[0-9]+}").HandlerFunc(h.Car.Update).Methods("PUT")
	v1.Path("/car/{id:[0-9]+}").HandlerFunc(h.Car.GetByID).Methods("GET")
	v1.Path("/car/{id:[0-9]+}").HandlerFunc(h.Car.Delete).Methods("DELETE")

	return router
}

type CarHandler interface {
	GetByID(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	Car CarHandler
}

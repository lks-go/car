package handler

import (
	"log"
	"net/http"
)

func responseOk(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(body)
}

func responseCreated(w http.ResponseWriter, location string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

func responseError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Write([]byte(http.StatusText(status)))

	if err != nil {
		log.Println("Error: ", err)
	}
}

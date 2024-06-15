package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	trips := TripResource{}

	mux.HandleFunc("GET /trips", trips.GetAll)
	mux.HandleFunc("POST /trips", trips.CreateOne)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Failed to listenn and serve: %v\n", err)
	}
}

type Trip struct { //логіка додатку
	ID          int
	Title       string
	Description string
	Source      string
	Destination string
}

type TripResource struct {
	s Storage
}

func (t *TripResource) GetAll(w http.ResponseWriter, r *http.Request) {
	trips := t.s.GetAllTrips()

	err := json.NewEncoder(w).Encode(trips)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (t *TripResource) CreateOne(w http.ResponseWriter, r *http.Request) {
	var trip Trip
	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		fmt.Printf("Failed to decode: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	trip.ID = t.s.CreateOneTrip(trip)
	err = json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		fmt.Printf("Failed to decode: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

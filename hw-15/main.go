package main

import (
	travelagency "hw15/internal/travel-agency"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	toursStorage := travelagency.NewInMemoryStorage()

	tourService := travelagency.NewService(toursStorage)

	toursHandler := travelagency.NewHandler(tourService)

	mux.HandleFunc("POST /tours", toursHandler.CreateTour)
	mux.HandleFunc("GET /tours", toursHandler.GetAllTours)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}

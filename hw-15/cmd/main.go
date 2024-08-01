package main

import (
	travelagency "hw15/internal/travel-agency"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	router := mux.NewRouter()

	toursStorage := travelagency.NewInMemoryStorage()
	idGenerator := travelagency.NewIDGenerator()
	tourService := travelagency.NewService(toursStorage, idGenerator)

	toursHandler := travelagency.NewHandler(tourService)

	toursHandler.RegisterRoutes(router)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}

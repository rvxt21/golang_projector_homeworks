package main

import (
	travelagency "hw15/internal/travel-agency"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	user := &travelagency.User{
		Name:  "Anastasiia",
		Email: "example@mail.com",
	}

	toursStorage := travelagency.NewInMemoryStorage()

	tourService := travelagency.NewService(toursStorage, user)

	toursHandler := travelagency.NewHandler(tourService)

	mux.HandleFunc("POST /tours", toursHandler.CreateTour)
	mux.HandleFunc("GET /tours", toursHandler.GetAllTours)
	mux.HandleFunc("/tours/book", toursHandler.BookTour) // Book a Tour
	mux.HandleFunc("/user/tours", toursHandler.GetUserTours)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}

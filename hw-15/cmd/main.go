package main

import (
	bookingservice "hw15/internal/booking-service"
	idgenerator "hw15/internal/id-generator"
	travelagency "hw15/internal/travel-agency"
	userservice "hw15/internal/user-service"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	router := mux.NewRouter()
	//Tours Service
	toursStorage := travelagency.NewInMemoryStorage()
	idGenerator := idgenerator.NewIDGenerator()
	tourService := travelagency.NewService(toursStorage, idGenerator)
	toursHandler := travelagency.NewHandler(tourService)

	toursHandler.RegisterRoutes(router)
	//User Service
	userStorage := userservice.NewInMemoryStorage()
	userService := userservice.NewUserService(&userStorage)
	userHandler := userservice.NewHandler(userService)
	userHandler.RegisterRoutes(router)

	//Reservation Service
	reservIdGenerator := idgenerator.NewIDGenerator()
	reservStorage := bookingservice.NewInMemStorage()
	reservService := bookingservice.NewService(&reservStorage, reservIdGenerator, userService, tourService)
	reservHandler := bookingservice.NewHandler(reservService)
	reservHandler.RegisterRoutes(router)

	//Server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}

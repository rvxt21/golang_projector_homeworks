package bookingservice

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type service interface {
	ReserveTour(tourID int, userID int)
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{s: s}
}

func (h Handler) CreateReservetion(w http.ResponseWriter, r *http.Request) {
	var op = "booking-service.handlers.CreateReservation"
	log.Info().Msgf("%s: making reservation", op)
}

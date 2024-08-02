package bookingservice

import (
	"encoding/json"
	"hw15/internal/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type service interface {
	ReserveTour(res Reservation) error
	GetReservationInfo(id int) (Reservation, bool)
	GetReservationsByUserID(userID int) []Reservation
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{s: s}
}

func (h Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/reservations", h.CreateReservetion).Methods("POST")
	r.Handle("/reservations/{id}", middlewares.IDHandler(http.HandlerFunc(h.GetReservationInfo))).Methods("GET")
	r.Handle("/users/{id}/myreservations", middlewares.IDHandler(http.HandlerFunc(h.GetReservationsByUserID))).Methods("GET")
}

func (h Handler) CreateReservetion(w http.ResponseWriter, r *http.Request) {
	var op = "booking-service.handlers.CreateReservation"
	log.Info().Msgf("%s: making reservation", op)

	var reserv Reservation

	err := json.NewDecoder(r.Body).Decode(&reserv)
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusInternalServerError)
		return
	}

	err = h.s.ReserveTour(reserv)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Handler) GetReservationInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middlewares.IdKey).(int)
	var reserv Reservation

	reserv, ok := h.s.GetReservationInfo(id)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewEncoder(w).Encode(reserv)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h Handler) GetReservationsByUserID(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middlewares.IdKey).(int)

	reservations := h.s.GetReservationsByUserID(id)

	err := json.NewEncoder(w).Encode(reservations)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

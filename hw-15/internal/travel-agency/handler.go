package travelagency

import (
	"encoding/json"
	"errors"
	"hw15/internal/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type service interface {
	CreateTour(tour Tour)
	GetAllTours() map[int]Tour
	GetTourByID(tourID int) (Tour, error)
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{s: s}
}

func (h Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tours", h.CreateTour).Methods("POST")
	r.HandleFunc("/tours", h.GetAllTours).Methods("GET")
	r.Handle("/tours/{id}", middlewares.IDHandler(http.HandlerFunc(h.GetTourInfo))).Methods("GET")
}

func (h Handler) CreateTour(w http.ResponseWriter, r *http.Request) {
	var tour Tour

	err := json.NewDecoder(r.Body).Decode(&tour)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Err(err).Msg("Error to decode JSON")
	}

	h.s.CreateTour(tour)
	w.WriteHeader(http.StatusCreated)

}

func (h Handler) GetTourInfo(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Get tour by ID Handler called")
	id := r.Context().Value(middlewares.IdKey).(int)

	tour, err := h.s.GetTourByID(id)
	if err != nil {
		if errors.Is(err, ErrTourNotFound) {
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, "Tour not found", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(tour)
	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h Handler) GetAllTours(w http.ResponseWriter, r *http.Request) {
	tours := h.s.GetAllTours()
	err := json.NewEncoder(w).Encode(tours)
	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON in GetAll")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

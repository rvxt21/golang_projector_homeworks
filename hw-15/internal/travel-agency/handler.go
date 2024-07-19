package travelagency

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type CreateTourReqBody struct {
	Title         string
	Price         uint16
	Programm      string
	TouristsNum   uint8 `json:"tourists_number"`
	Nutrition     Nutrition
	TransportType Transport `json:"transport_type"`
}
type service interface {
	CreateTour(title string, price uint16, programm string, touristsnum uint8, nutrition Nutrition, transport Transport)
	GetAll() []Tour
	GetUserTours() []Tour
	BookTour(id string)
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{s: s}
}

func (h Handler) CreateTour(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateTourReqBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Err(err).Msg("Error to decode JSON")
	}

	h.s.CreateTour(reqBody.Title, reqBody.Price, reqBody.Programm, reqBody.TouristsNum, reqBody.Nutrition, reqBody.TransportType)

}

func (h Handler) GetAllTours(w http.ResponseWriter, r *http.Request) {
	tours := h.s.GetAll()
	err := json.NewEncoder(w).Encode(tours)
	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON in GetAll")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h Handler) BookTour(w http.ResponseWriter, r *http.Request) {
	tourID := r.URL.Query().Get("id")
	if tourID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Msg("Missing tour ID")
		return
	}

	h.s.BookTour(tourID)
	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetUserTours(w http.ResponseWriter, r *http.Request) {
	tours := h.s.GetUserTours()
	err := json.NewEncoder(w).Encode(tours)
	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON in GetUserTours")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

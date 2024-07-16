package travelagency

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func CreateTour(w http.ResponseWriter, r *http.Request) {
	var t Tour

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Err(err).Msg("Error to decode JSON")
	}

}

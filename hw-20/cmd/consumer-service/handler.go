package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func OrangesAnalyticsHandler(s *OrangeService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(s.GetAnalytics())
		if err != nil {
			log.Warn().Err(err).Msg("Failed to JSON encode")
		}
	})
}

func AllOrangesHandler(s *OrangeService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(s.Storage.GetAllOranges())
		if err != nil {
			log.Warn().Err(err).Msg("Failed to JSON encode")
		}
	})
}

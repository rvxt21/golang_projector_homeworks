package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func CallHandler(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msgf("error to get response")
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Error().Msgf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Info().Err(err)
	}
	fmt.Printf("%s", body)
}

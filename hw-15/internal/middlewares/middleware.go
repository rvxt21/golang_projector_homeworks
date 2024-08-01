package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type ContextKey string

const IdKey ContextKey = "id"

func IDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Middleware called")
		vars := mux.Vars(r)
		IDStr, ok := vars["id"]
		if !ok {
			log.Info().Msg("Missed ID in request")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ID, err := strconv.Atoi(IDStr)
		if err != nil {
			log.Warn().Err(err).Msg("Invalid ID param")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), IdKey, ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

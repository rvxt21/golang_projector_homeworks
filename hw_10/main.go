package main

import (
	"hw10/resources"
	"hw10/storage"
	"net/http"

	"hw10/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	mux := mux.NewRouter()

	taskRes := resources.TasksResourse{
		S: storage.NewStorage(),
	}

	mux.HandleFunc("/task", taskRes.CreateTask).Methods("POST")
	mux.HandleFunc("/tasks", taskRes.GetAll).Methods("GET")
	mux.Handle("/tasks/{id}", middleware.IdMiddleware(http.HandlerFunc(taskRes.DeleteTask))).Methods("DELETE")
	mux.Handle("/tasks/{id}", middleware.IdMiddleware(http.HandlerFunc(taskRes.UpdateTask))).Methods("PUT")
	mux.HandleFunc("/user", taskRes.CreateUser).Methods("POST")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}

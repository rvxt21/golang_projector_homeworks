package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	taskRes := TasksResourse{
		s: NewStorage(),
	}

	mux.HandleFunc("POST /task", taskRes.CreateOne)
	mux.HandleFunc("GET /tasks", taskRes.GetAll)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}

type TasksResourse struct {
	s *Storage
}

func (tr *TasksResourse) CreateOne(w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Fatal().Err(err).Msg("Error to decode JSON in CreateOneTasks")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task.ID = tr.s.CreateOneTask(task)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Fatal().Err(err).Msg("Error to encode JSON in CreateOneTasks")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (tr *TasksResourse) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks := tr.s.GetAllTasks()
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Fatal().Err(err).Msg("Error to encode JSON in GetAll")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

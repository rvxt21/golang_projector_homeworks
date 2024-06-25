package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	taskRes := TasksResourse{
		s: NewStorage(),
	}

	mux.HandleFunc("POST /task", taskRes.CreateTask)
	mux.HandleFunc("GET /tasks", taskRes.GetAll)
	mux.HandleFunc("DELETE /task/{id}", taskRes.DeleteTask)
	mux.HandleFunc("PUT /task/{id}", taskRes.UpdateTask)
	mux.HandleFunc("POST /user", taskRes.CreateUser)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}

type TasksResourse struct {
	s *Storage
}

func (tr *TasksResourse) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Error().Err(err).Msg("Error to decode JSON in CreateOneupdates")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = tr.s.CreateOneTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(task)

	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON in CreateOneupdates")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (tr *TasksResourse) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks := tr.s.GetAllTasks()
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON in GetAll")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (tr *TasksResourse) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idVal := r.PathValue("id")

	taskID, err := strconv.Atoi(idVal)
	if err != nil {
		log.Warn().Err(err).Msg("Invalid ID param")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := tr.s.DeleteTask(taskID)
	if !ok {
		log.Info().Msg("Task to delete not found.")
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func (tr *TasksResourse) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idVal := r.PathValue("id")

	updateID, err := strconv.Atoi(idVal)
	if err != nil {
		log.Error().Err(err).Msg("Invalid ID param")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var update Task
	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Error().Err(err).Msg("Error to decode JSON in UpdateTask")
	}

	updated := tr.s.UpdateTask(updateID, update)
	if !updated {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(update)
}

func (tr TasksResourse) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Error().Err(err).Msg("Error to decode JSON in CreateUser")
		return
	}
	_, err = tr.s.CreateOneUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON in CreateOneupdates")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

package resources

import (
	"encoding/json"
	"hw10/enteties"
	"hw10/middleware"
	"hw10/storage"
	"net/http"

	"github.com/rs/zerolog/log"
)

type TasksResourse struct {
	S *storage.Database
}

func (tr *TasksResourse) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task enteties.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Error().Err(err).Msg("Error to decode JSON in CreateOneupdates")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tr.S.InsertTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(task)

	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON in CreateOneTask")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (tr *TasksResourse) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := tr.S.GetAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		log.Error().Err(err).Msgf("Error getting all tasks")
	}
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON in GetAll")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (tr *TasksResourse) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.Context().Value(middleware.IdKey).(int)

	ok, err := tr.S.DeleteTask(taskID)
	if err != nil {
		if err == storage.ErrDeleteFailed {
			log.Error().Err(err).Msgf("unable to delete task")
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		if !ok {
			log.Error().Err(err).Msgf("unable to find task to delete")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (tr *TasksResourse) GetTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.Context().Value(middleware.IdKey).(int)
	tasks, _, err := tr.S.GetTaskByID(taskID)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		log.Error().Err(err).Msgf("Error getting task")
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Error().Err(err).Msg("Error to encode JSON in GetAll")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// func (tr *TasksResourse) UpdateTask(w http.ResponseWriter, r *http.Request) {
// 	updateID := r.Context().Value(middleware.IdKey).(int)

// 	var update enteties.Task
// 	err := json.NewDecoder(r.Body).Decode(&update)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Error to decode JSON in UpdateTask")
// 	}

// 	updated := tr.S.UpdateTask(updateID, update)
// 	if !updated {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(update)
// }

// func (tr TasksResourse) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var user enteties.User

// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Error to decode JSON in CreateUser")
// 		return
// 	}
// 	_, err = tr.S.CreateOneUser(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(user)

// 	if err != nil {
// 		log.Error().Err(err).Msg("Error to encode JSON in CreateOneupdates")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// }

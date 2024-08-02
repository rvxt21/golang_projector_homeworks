package userservice

import (
	"encoding/json"
	"hw15/internal/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type service interface {
	CreateUser(u User)
	GetUserById(id int) (User, bool)
}

type Handler struct {
	serv service
}

func NewHandler(s service) Handler {
	return Handler{serv: s}
}

func (h Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
	// r.HandleFunc("/tours", h.GetAllUsers).Methods("GET")
	r.Handle("/users/{id}", middlewares.IDHandler(http.HandlerFunc(h.GetUserInfoById))).Methods("GET")
}
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusInternalServerError)
		return
	}

	h.serv.CreateUser(user)
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) GetUserInfoById(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middlewares.IdKey).(int)
	user, ok := h.serv.GetUserById(id)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

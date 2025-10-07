package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Maria"},
	{ID: 2, Name: "Jose"},
}

// GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	for _, u := range users {
		if idParam == string(rune('0'+u.ID)) {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

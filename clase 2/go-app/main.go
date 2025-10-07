package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-app/handler"
)

func main() {
	r := chi.NewRouter()

	// Middlewares útiles
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rutas
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("👋 Bienvenido a la API!"))
	})

	// Rutas de usuario
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
		r.Get("/{id}", handler.GetUserByID)
	})

	http.ListenAndServe(":8089", r)
}

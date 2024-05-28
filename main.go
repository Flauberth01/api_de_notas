package main

import (
	"net/http"

	"api_de_notas/database"
	"api_de_notas/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database.ConectaBanco()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/notas", handlers.ExibeNotas)

	http.ListenAndServe(":8000", r)
}

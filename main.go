package main

import (
	"net/http"

	"api_de_notas/database"
	"api_de_notas/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "api_de_notas/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	database.ConectaBanco()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/notas/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/notas/swagger/doc.json"),
	))

	r.Get("/notas", handlers.ExibeNotas)
	r.Post("/notas", handlers.CriaNota)
	r.Put("/notas/{id}", handlers.AtualizaNota)
	r.Delete("/notas/{id}", handlers.DeletaNota)

	http.ListenAndServe(":8000", r)
}

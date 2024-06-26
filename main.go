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

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	database.ConectaBanco()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(jsonMiddleware)

	r.Get("/notas/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/notas/swagger/doc.json"),
	))

	r.Get("/notas", handlers.ExibeNotas)
	r.Post("/notas", handlers.CriaNota)
	r.Put("/notas/{id}", handlers.AtualizaNota)
	r.Delete("/notas/{id}", handlers.DeletaNota)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
}

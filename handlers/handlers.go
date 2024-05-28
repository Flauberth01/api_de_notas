package handlers

import (
	"encoding/json"
	"net/http"

	"api_de_notas/services"
)

func ExibeNotas(w http.ResponseWriter, r *http.Request) {
	notas, err := services.ExibeNotas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notas)
}

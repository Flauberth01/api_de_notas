package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api_de_notas/models"
	"api_de_notas/services"
	"github.com/go-chi/chi/v5"
)

type CriaNotasPayload struct {
	Titulo   string `json:"titulo"`
	Conteudo string `json:"conteudo"`
}

// ExibeNotas @Summary Exibe as notas
// @Description Exibe todas as notas
// @Tags notas
// @Accept json
// @Produce json
// @Success 200 {array} models.Nota
// @Router /notas [get]
func ExibeNotas(w http.ResponseWriter, r *http.Request) {
	notas, err := services.ExibeNotas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notas)
}

// CriaNota @Summary Cria uma nota
// @Description Cria uma nova nota
// @Tags notas
// @Accept json
// @Produce json
// @Param payload body CriaNotasPayload true "Payload para criar uma nova nota"
// @Success 201
// @Router /notas [post]
func CriaNota(w http.ResponseWriter, r *http.Request) {
	var notasPayload CriaNotasPayload
	err := json.NewDecoder(r.Body).Decode(&notasPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nota := models.Nota{
		Titulo:   notasPayload.Titulo,
		Conteudo: notasPayload.Conteudo,
	}

	err = services.CriaNota(nota)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// AtualizaNota @Summary Atualiza uma nota
// @Description Atualiza uma nota existente
// @Tags notas
// @Accept json
// @Produce json
// @Param id path string true "ID da nota"
// @Param payload body CriaNotasPayload true "Payload para atualizar uma nova nota"
// @Success 204
// @Router /notas/{id} [put]
func AtualizaNota(w http.ResponseWriter, r *http.Request) {
	notaID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var notasPayload CriaNotasPayload

	err = json.NewDecoder(r.Body).Decode(&notasPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nota := models.Nota{
		ID:       notaID,
		Titulo:   notasPayload.Titulo,
		Conteudo: notasPayload.Conteudo,
	}

	err = services.AtualizaNota(nota)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeletaNota @Summary Deleta uma nota
// @Description Deleta uma nota existente
// @Tags notas
// @Accept json
// @Produce json
// @Param id path string true "ID da nota"
// @Success 204
// @Router /notas/{id} [delete]
func DeletaNota(w http.ResponseWriter, r *http.Request) {
	notaID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.DeletaNota(notaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

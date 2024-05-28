package services

import (
	"api_de_notas/database"
	"api_de_notas/models"
)

func ExibeNotas() ([]models.Nota, error) {
	var notas []models.Nota
	result := database.DB.Find(&notas)
	return notas, result.Error
}

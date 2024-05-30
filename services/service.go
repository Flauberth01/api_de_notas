package services

import (
	"errors"

	"api_de_notas/database"
	"api_de_notas/models"
)

func ExibeNotas() ([]models.Nota, error) {
	var notas []models.Nota
	err := database.DB.Find(&notas).Error
	return notas, err
}

func CriaNota(nota models.Nota) error {
	err := database.DB.Create(&nota).Error
	return err
}

func AtualizaNota(nota models.Nota) error {
	var notaExistente models.Nota

	if err := database.DB.First(&notaExistente, nota.ID).Error; err != nil {
		return errors.New("ID não encontrado")
	}

	err := database.DB.Model(&notaExistente).Updates(nota).Error

	return err
}

func DeletaNota(id uint64) error {
	var nota models.Nota

	if err := database.DB.First(&nota, id).Error; err != nil {
		return errors.New("ID não encontrado")
	}

	err := database.DB.Delete(&nota).Error

	return err
}

package services

import (
	"http_fiber_api/app/models"

	"gorm.io/gorm"
)

func CreateRecurso(db *gorm.DB, recurso *models.Recurso) error {
	result := db.Create(&recurso) // Passa a referência do modelo
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllRecursos(db *gorm.DB, recursos *[]models.Recurso) error {
	result := db.Find(&recursos) // Passa a referência da slice
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetRecursoByID(db *gorm.DB, recurso *models.Recurso, id uint64) error {
	result := db.First(&recurso, id) // Passa a referência do modelo e o ID
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateRecurso(db *gorm.DB, recurso *models.Recurso) error {
	result := db.Save(&recurso) // Passa a referência do modelo atualizado
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteRecurso(db *gorm.DB, recurso *models.Recurso) error {
	result := db.Delete(&recurso) // Passa a referência do modelo
	if result.Error != nil {
		return result.Error
	}
	return nil
}

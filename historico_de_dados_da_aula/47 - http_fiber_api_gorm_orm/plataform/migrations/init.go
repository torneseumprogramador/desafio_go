package migrations

import (
	"http_fiber_api/app/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Recurso{})
}

package migrations

import (
	"fmt"
	"http_fiber_api/app/models"
	"http_fiber_api/plataform/database"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Administrador{})
	db.AutoMigrate(&models.Recurso{})

	fmt.Println("==== Db Migrate com sucesso ====")
}

func Seed() {
	db := database.GetConnection()

	var count int64
	email := "danilo@teste.com"

	db.Model(&models.Administrador{}).Where("email = ?", email).Count(&count)

	if count == 0 {
		adm := models.Administrador{
			Nome:  "Danilo",
			Email: email,
			Senha: "123456",
		}

		result := db.Create(&adm)

		if result.Error != nil {
			fmt.Println("======= Erro ao rodar seed ======")
			fmt.Println(result.Error)
			return
		}
	}

	fmt.Println("=== Seed executada com sucesso ===")
}

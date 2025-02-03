package migrations

import (
	"amplassignment/database"
	"amplassignment/models"
)

func Migrate() {
	database.DB.AutoMigrate(&models.Task{})
}

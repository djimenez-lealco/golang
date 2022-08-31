package migrations

import (
	"orm/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Client{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Orderitem{})
}

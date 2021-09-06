package migrations

import (
	"github.com/ueverson/ProcessingWorksheetGO/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Asset{})
}

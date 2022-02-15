package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1631657069CreateProgressTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1631657069",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Progress{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Model(&models.Progress{}).DropTable("progresses").Error
		},
	}
}

//Get-Date -UFormat %s

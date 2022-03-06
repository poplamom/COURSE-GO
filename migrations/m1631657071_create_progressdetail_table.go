package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1631657071CreateProgressDetailTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1631657071",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.ProgressDetail{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Model(&models.ProgressDetail{}).DropTable("progressdetails").Error
		},
	}
}

//Get-Date -UFormat %s

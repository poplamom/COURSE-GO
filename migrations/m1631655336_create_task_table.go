package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1631655336CreateTaskTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1631655336",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Task{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("users").Error
		},
	}
}
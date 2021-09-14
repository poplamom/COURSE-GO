package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1631656943AddTaskIDToCourses() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1631656943",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Course{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Model(&models.Course{}).DropColumn("task_id").Error
		},
	}
}

//Get-Date -UFormat %s
package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1631657068AddQuestionIDToTasks() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1631657068",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Task{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Model(&models.Task{}).DropColumn("question_id").Error
		},
	}
}

//Get-Date -UFormat %s
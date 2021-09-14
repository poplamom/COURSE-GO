package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1631656197AddUserIDToUserCourses() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1631656197",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.UserCourse{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Model(&models.UserCourse{}).DropColumn("user_id").Error
		},
	}
}

//Get-Date -UFormat %s
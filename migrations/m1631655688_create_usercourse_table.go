package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1631655688CreateUserCourseTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1631655688",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.UserCourse{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("users").Error
		},
	}
}
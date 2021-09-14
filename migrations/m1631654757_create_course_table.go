package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1631654757CreateCourseTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1631654757",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Course{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("users").Error
		},
	}
}
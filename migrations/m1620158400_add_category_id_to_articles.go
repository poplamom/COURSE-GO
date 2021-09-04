package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1620158400AddCategoryIDToArticles() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1620158400",
		Migrate: func(tx *gorm.DB) error {
			err := tx.AutoMigrate(&models.Article{}).Error

			var articles []models.Article
			tx.Unscoped().Find(&articles)
			for _ , ararticles := range articles{
				ararticles.CategoryID = 2
				tx.Save(&ararticles)
			}

			return err
		},
		Rollback: func(tx *gorm.DB) error{
			return tx.Model(&models.Article{}).DropColumn("category_id").Error
			
		},
	}
}

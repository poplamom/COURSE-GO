package migrations

import (
	"course-go/config"
	"log"

	"gopkg.in/gormigrate.v1"
)

func Migrate() {
db := config.GetDB()
m := gormigrate.New(
	db,
	gormigrate.DefaultOptions,
	[]*gormigrate.Migration{
		m1618442592CreateArticlesTable(),
		m1620154439CreateCategoriesTable(),
		m1620158400AddCategoryIDToArticles(),
		m1620165273CreateUsersTable(),
		m1620169040AddUserIDToArticles(),
	},
)

if err := m.Migrate(); err != nil {
	log.Fatalf("Could not migrate: %v", err)
}
}

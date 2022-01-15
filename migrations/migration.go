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
		m1631654757CreateCourseTable(),
		m1631655336CreateTaskTable(),
		m1631655506CreateQuestionTable(),
		m1631655688CreateUserCourseTable(),
		// m1631656197AddUserIDToUserCourses(),
		// m1631656721AddCourseIDToUserCourses(),
		// m1631656943AddTaskIDToCourses(),
		// m1631657068AddQuestionIDToTasks(),
	},
)

if err := m.Migrate(); err != nil {
	log.Fatalf("Could not migrate: %v", err)
}
}

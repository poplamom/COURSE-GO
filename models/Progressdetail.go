package models

import "github.com/jinzhu/gorm"

type ProgressDetail struct {
	gorm.Model
	CourseID   uint
	TaskID     uint
	QuestionID uint
	UserID     uint
	Course     Course
	Task       Task
	Question   Question
	User       User
}

package models

import "github.com/jinzhu/gorm"

type ProgressDetail struct {
	gorm.Model
	TaskID     uint
	QuestionID uint
	UserID     uint
	Task       Task
	Question   Question
	User       User
}

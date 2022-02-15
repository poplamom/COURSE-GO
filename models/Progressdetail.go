package models

import "github.com/jinzhu/gorm"

type ProgressDetail struct {
	gorm.Model
	ProgressID uint
	TaskID     uint
	QuestionID uint
	Task       Task
	Progress   Progress
	Question   Question
}

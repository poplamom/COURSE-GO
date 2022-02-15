package models

import "github.com/jinzhu/gorm"

type Progress struct {
	gorm.Model
	CourseID       uint
	Course         Course
	ProgressDetail []ProgressDetail
}

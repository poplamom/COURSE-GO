package models

import "github.com/jinzhu/gorm"

type Progress struct {
	gorm.Model
	UserID   uint
	CourseID uint
	Course   Course
	User     User
}

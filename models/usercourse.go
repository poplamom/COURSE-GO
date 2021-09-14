package models

import "github.com/jinzhu/gorm"

type UserCourse struct {
	gorm.Model
	UserID uint
	CourseID uint
	Progress	int `gorm:"not null"`
}
package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Desc      string `gorm:"not null"`
	Objective string `gorm:"not null"`
	Status    string `gorm:"default:'0'"`
	CourseID  uint
	Course    Course
	Question  []Question
}

package models

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name      	string `gorm:"unique;not null"`
	Desc 		string `gorm:"not null"`
	TaskID uint
}
package models

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Name		string `gorm:"unique;not null"`
	Answer		string `gorm:"unique;not null"`
	Hint		string `gorm:"unique;not null"`
	Status 		int `gorm:"unique;not null"`
}
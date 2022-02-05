package models

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Answer string `gorm:"not null"`
	Hint   string `gorm:"not null"`
	TaskID uint
	Task   Task
}

package models

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Title       string
	Description string
	Finished    bool `gorm:"default:false"`
	UserID      uint
}

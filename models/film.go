package models

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	KodeFilm string `gorm:"type:varchar(10);unique;not null"`
	Name     string `grom:"type:varchar(100);not null"`
	Genre    string `gorm:"type:varchar(100)"`
	Year     int    `gorm:"not null"`
}

package models

import (
	"gorm.io/gorm"
)

type Kategori struct {
	gorm.Model
	Name    string    `gorm:"type:varchar(100);not null"`
	Product []Product `gorm:"foreignKey:KategoriID"`
}

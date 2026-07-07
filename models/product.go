package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string   `gorm:"type:varchar(100);not null"`
	Price      uint     `gorm:"type:int;not null"`
	KategoriID uint     `gorm:"type:not null"`
	Kategori   Kategori `gorm:"foreignKey:KategoriID;constraint:OnUpdate:CASCADE:RESTRICT;"`
}

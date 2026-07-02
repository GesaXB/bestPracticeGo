package models

type Book struct {
	Id       uint   `gorm:"type:int;primaryKey"`
	Nama     string `gorm:"type:varchar(100)" json:"nama"`
	Author   string `gorm:"type:varchar(100)" json:"author"`
	Kategori string `gorm:"type:varchar(100)" json:"kategori"`
}

package models

type Makanan struct {
	Id          uint   `gorm:"type:int;primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:TEXT"`
}

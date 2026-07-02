package models

type Makanan struct {
	Id          uint   `gorm:"type:int;primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:TEXT" json:"description"`
}

func (Makanan) TableName() string {
	return "makanan"
}

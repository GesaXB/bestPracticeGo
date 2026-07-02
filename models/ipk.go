package models

type IPK struct {
	Nim      uint    `gorm:"type:int" json:"nin"`
	Semseter string  `gorm:"type:varchar(3)" json:"semseter"`
	Ipk      float64 `grom:"type:decimal(10,2)" json:"ipk"`
}

func (IPK) TableName() string {
	return "ipk"
}

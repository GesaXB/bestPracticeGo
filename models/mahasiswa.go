package models

type Mahasiswa struct {
	Nim          uint   `gorm:"primaryKey" json:"nim"`
	NamaDepan    string `gorm:"type:varchar(50)" json:"nama_depan"`
	NamaBelakang string `gorm:"type:varchar(50)" json:"nama_belakang"`
	KotaAsal     string // default jadi TEXT
	Jurusan      string
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}

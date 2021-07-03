package models

import (
	"labpro-backend/pkg/config"

	"gorm.io/gorm"
)

type Dorayaki struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	Rasa      string
	Deskripsi string
	Gambar    string
}

func (b *Dorayaki) CreateDorayaki() *Dorayaki {
	config.DB.Create(&b)
	return b
}

func GetAllDorayakis() []Dorayaki {
	var Dorayakis []Dorayaki
	config.DB.Find(&Dorayakis)
	return Dorayakis
}

func GetDorayakiById(Id int64) (*Dorayaki, *gorm.DB) {
	var getDorayaki Dorayaki
	db := config.DB.Where("ID = ?", Id).Find(&getDorayaki)
	return &getDorayaki, db
}

func DeleteDorayaki(ID int64) Dorayaki {
	var Dorayaki Dorayaki
	config.DB.Where("ID = ?", ID).Delete(Dorayaki)
	return Dorayaki
}

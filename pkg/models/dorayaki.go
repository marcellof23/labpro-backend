package models

import (
	"labpro-backend/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Dorayaki struct {
	gorm.Model
	//Id          string `json:"id"`
	Rasa      string `gorm:""json:"rasa"`
	Deskripsi string `json:"deskripsi"`
	Gambar    string `json:"gambar"`
}

func init() {
	if config.Connect() != nil {
		panic(config.Connect())
	}

	db = config.GetDB()
	db.AutoMigrate(&Dorayaki{})
}

func (b *Dorayaki) CreateDorayaki() *Dorayaki {
	db.Create(&b)
	return b
}

func GetAllDorayakis() []Dorayaki {
	var Dorayakis []Dorayaki
	db.Find(&Dorayakis)
	return Dorayakis
}

func GetDorayakiById(Id int64) (*Dorayaki, *gorm.DB) {
	var getDorayaki Dorayaki
	db := db.Where("ID = ?", Id).Find(&getDorayaki)
	return &getDorayaki, db
}

func DeleteDorayaki(ID int64) Dorayaki {
	var Dorayaki Dorayaki
	db.Where("ID = ?", ID).Delete(Dorayaki)
	return Dorayaki
}

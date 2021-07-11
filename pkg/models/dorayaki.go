package models

import (
	"fmt"
	"labpro-backend/pkg/config"
	"time"

	"gorm.io/gorm"
)

type Dorayaki struct {
	ID              int64 `gorm:"primaryKey"`
	Rasa            *string
	Deskripsi       *string
	Gambar          *string
	DorayakiStoreID *int64
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at"`
}

func (Dorayaki) TableName() string {
	return "dorayaki"
}

func (b *Dorayaki) CreateDorayaki() *Dorayaki {
	config.DB.Create(&b)
	return b
}

func GetAllDorayakis() []Dorayaki {
	var Dorayakis []Dorayaki
	config.DB.Find(&Dorayakis)
	fmt.Println("Hello")
	fmt.Println(Dorayakis)
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

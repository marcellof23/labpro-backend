package models

import (
	"errors"
	"labpro-backend/pkg/config"
	"time"

	"gorm.io/gorm"
)

type Dorayaki struct {
	ID              int64 `gorm:"primaryKey"`
	Rasa            *string
	Deskripsi       *string
	Gambar          *string
	Jumlah          *int64
	DorayakiStoreID *int64
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at"`
}

func (Dorayaki) TableName() string {
	return "dorayaki"
}

func (b *Dorayaki) CreateDorayaki() (*Dorayaki, error) {
	var DistinctVariant []Dorayaki
	var foundVariant bool
	config.DB.Distinct("rasa").Find(&DistinctVariant)

	for _, v := range DistinctVariant {
		if *v.Rasa == *b.Rasa {
			foundVariant = true
		}
	}

	var err error
	if !foundVariant {
		config.DB.Create(&b)
	} else {
		err = errors.New("rasa is already provided in this store")
	}
	return b, err
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

func GetDorayakiByIdStore(Id int64) ([]Dorayaki, *gorm.DB) {
	var getDorayakiList []Dorayaki
	db := config.DB.Where("dorayaki_store_id = ?", Id).Find(&getDorayakiList)
	return getDorayakiList, db
}

func GetAllVariant() []Dorayaki {
	var DorayakisVariant []Dorayaki
	config.DB.Distinct("rasa").Find(&DorayakisVariant)
	return DorayakisVariant
}

func GetVariantById(Id int64) ([]Dorayaki, *gorm.DB) {
	var getDorayakiVariant []Dorayaki
	db := config.DB.Distinct("rasa").Where("dorayaki_store_id = ?", Id).Find(&getDorayakiVariant)
	return getDorayakiVariant, db
}

func DeleteDorayaki(ID int64) Dorayaki {
	var Dorayaki Dorayaki
	config.DB.Where("ID = ?", ID).Delete(Dorayaki)
	return Dorayaki
}

package models

import (
	"labpro-backend/pkg/config"

	"gorm.io/gorm"
)

type DorayakiStore struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	Nama      string
	Jalan     string
	Kecamatan string
	Provinsi  string
}

func (b *DorayakiStore) CreateDorayakiStore() *DorayakiStore {
	config.DB.Create(&b)
	return b
}

func GetAllDorayakiStores() []DorayakiStore {
	var DorayakiStores []DorayakiStore
	config.DB.Find(&DorayakiStores)
	return DorayakiStores
}

func GetDorayakiStoreById(Id int64) (*DorayakiStore, *gorm.DB) {
	var getDorayakiStore DorayakiStore
	db := config.DB.Where("ID = ?", Id).Find(&getDorayakiStore)
	return &getDorayakiStore, db
}

func DeleteDorayakiStore(ID int64) DorayakiStore {
	var DorayakiStore DorayakiStore
	config.DB.Where("ID = ?", ID).Delete(DorayakiStore)
	return DorayakiStore
}

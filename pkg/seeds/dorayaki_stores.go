package seeds

import (
	"labpro-backend/pkg/models"

	"gorm.io/gorm"
)

func CreateDorayakiStores(db *gorm.DB, nama, jalan, kecamatan, provinsi string) error {
	return db.Create(&models.DorayakiStore{Nama: nama, Jalan: jalan, Kecamatan: kecamatan, Provinsi: provinsi}).Error
}

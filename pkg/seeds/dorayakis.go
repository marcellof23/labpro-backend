package seeds

import (
	"labpro-backend/pkg/models"

	"gorm.io/gorm"
)

func CreateDorayaki(db *gorm.DB, rasa, deskripsi, gambar string) error {
	return db.Create(&models.Dorayaki{Rasa: rasa, Deskripsi: deskripsi, Gambar: gambar}).Error
}

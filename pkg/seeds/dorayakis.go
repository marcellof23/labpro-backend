package seeds

import (
	"labpro-backend/pkg/models"

	"gorm.io/gorm"
)

func CreateDorayaki(db *gorm.DB, rasa, deskripsi, gambar *string, DorayakiStoreID *int64) error {
	return db.Create(&models.Dorayaki{Rasa: rasa, Deskripsi: deskripsi, Gambar: gambar, DorayakiStoreID: DorayakiStoreID}).Error
}

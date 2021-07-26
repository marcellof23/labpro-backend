package seeds

import (
	"labpro-backend/pkg/seed"

	"gorm.io/gorm"
)

func AllDorayaki() []seed.Seed {
	return []seed.Seed{
		{
			Name: "CreateJane",
			Run: func(db *gorm.DB) error {
				var rasa = "Coklat"
				var deskripsi = "enak"
				var gambar = "gada"
				var jumlah = int64(3)
				var dorayakistoreId = int64(1)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &jumlah, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil

			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "Strawberry"
				var deskripsi = "enak"
				var gambar = "gada"
				var jumlah = int64(5)
				var dorayakistoreId = int64(1)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &jumlah, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "Keju"
				var deskripsi = "enak"
				var gambar = "gada"
				var jumlah = int64(2)
				var dorayakistoreId = int64(1)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &jumlah, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "Mangga"
				var deskripsi = "enak"
				var gambar = "gada"
				var jumlah = int64(5)
				var dorayakistoreId = int64(2)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &jumlah, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "Apel"
				var deskripsi = "enak"
				var gambar = "gada"
				var jumlah = int64(3)
				var dorayakistoreId = int64(2)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &jumlah, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "Coklat"
				var deskripsi = "Ngga enak"
				var gambar = "gada"
				var jumlah = int64(4)
				var dorayakistoreId = int64(3)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &jumlah, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateAssdf",
			Run: func(db *gorm.DB) error {
				var rasa = "Jeruk"
				var deskripsi = "enak"
				var gambar = "gada"
				var jumlah = int64(3)
				var dorayakistoreId = int64(4)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &jumlah, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateAssdf",
			Run: func(db *gorm.DB) error {
				var rasa = "Apel"
				var deskripsi = "enak"
				var gambar = "gada"
				var jumlah = int64(4)
				var dorayakistoreId = int64(5)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &jumlah, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}

func AllDorayakiStore() []seed.Seed {
	return []seed.Seed{
		{
			Name: "Toko1",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko1", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil

			},
		},
		{
			Name: "Toko2",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko2", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "Toko3",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko3", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "Toko4",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko4", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "Toko5",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko5", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}

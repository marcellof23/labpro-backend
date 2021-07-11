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
				var rasa = "jane"
				var deskripsi = "enak"
				var gambar = "gada"
				var dorayakistoreId = int64(2)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil

			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "John"
				var deskripsi = "enak"
				var gambar = "gada"
				var dorayakistoreId = int64(2)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "Yow"
				var deskripsi = "enak"
				var gambar = "gada"
				var dorayakistoreId = int64(2)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "Hehe"
				var deskripsi = "enak"
				var gambar = "gada"
				var dorayakistoreId = int64(3)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "Lolo"
				var deskripsi = "enak"
				var gambar = "gada"
				var dorayakistoreId = int64(2)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				var rasa = "Fafa"
				var deskripsi = "Ngga enak"
				var gambar = "gada"
				var dorayakistoreId = int64(1)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &dorayakistoreId)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateAssdf",
			Run: func(db *gorm.DB) error {
				var rasa = "Asdf"
				var deskripsi = "enak"
				var gambar = "gada"
				var dorayakistoreId = int64(1)
				err := CreateDorayaki(db, &rasa, &deskripsi, &gambar, &dorayakistoreId)
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
			Name: "CreateJane",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko14", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil

			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko13", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko21", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko2", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko9", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko8", "Jalan1", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateAssdf",
			Run: func(db *gorm.DB) error {
				err := CreateDorayakiStores(db, "Toko4", "Jalan5", "KecamatanA", "ProvinsiA")
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}

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
				err := CreateDorayaki(db, "Jane", "enak", "gada", 2)
				if err != nil {
					return err
				}
				return nil

			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "John", "enak", "gada", 3)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Yow", "enak", "gada", 2)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Hehe", "enak", "gada", 3)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Lolo", "enak", "gada", 2)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Fafa", "enak", "gada", 1)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateAssdf",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Asdf", "enak", "gada", 1)
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

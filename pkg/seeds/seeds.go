package seeds

import (
	"labpro-backend/pkg/seed"

	"gorm.io/gorm"
)

func All1() []seed.Seed {
	return []seed.Seed{
		{
			Name: "CreateJane",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Jane", "enak", "gada")
				if err != nil {
					return err
				}
				return nil

			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "John", "enak", "gada")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Yow", "enak", "gada")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Hehe", "enak", "gada")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Lolo", "enak", "gada")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Fafa", "enak", "gada")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateAssdf",
			Run: func(db *gorm.DB) error {
				err := CreateDorayaki(db, "Asdf", "enak", "gada")
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}

func All2() []seed.Seed {
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

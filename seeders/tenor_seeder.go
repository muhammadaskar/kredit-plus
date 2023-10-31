package seeders

import (
	"log"

	"github.com/muhammadaskar/kredit-plus/domains"
	"gorm.io/gorm"
)

func SeedTenors(db *gorm.DB) {
	tenors := []domains.Tenor{
		{
			ID:          "tenor-0c8215f2-bcb8-4b2b-bb73-6558c6d9a64d",
			ConsumerID:  "user-6af79619-c66f-4df7-98a2-b0ee282a3cc1",
			LimitTenor1: 100000,
			LimitTenor2: 200000,
			LimitTenor3: 500000,
			LimitTenor4: 700000,
		},
		{
			ID:          "tenor-34d452b0-c4b2-43b5-9db6-45e73592cd34",
			ConsumerID:  "user-e17227f6-19d5-4403-a12b-be8a77c74897",
			LimitTenor1: 1000000,
			LimitTenor2: 1200000,
			LimitTenor3: 1500000,
			LimitTenor4: 2000000,
		},
	}

	for _, t := range tenors {
		result := db.Create(&t)
		if result.Error != nil {
			log.Println("Error seeding tenor data:", result.Error)
			return
		}
	}
}

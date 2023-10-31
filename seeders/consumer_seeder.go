package seeders

import (
	"log"

	"github.com/muhammadaskar/kredit-plus/domains"
	customeparse "github.com/muhammadaskar/kredit-plus/utils/custome_parse"
	"gorm.io/gorm"
)

func SeedConsumers(db *gorm.DB) {
	consumers := []domains.Consumer{
		{
			ID:           "user-6af79619-c66f-4df7-98a2-b0ee282a3cc1",
			NIK:          "fc1742154cd0fc06ba378ad35886b478735c8f37fe7d02fdaaa62f0bb4c93c13",
			FullName:     "Budi",
			LegalName:    "Budi Sutejo",
			TempatLahir:  "Jakarta",
			TanggalLahir: customeparse.ParseTime("2000-01-01"),
			Gaji:         1500000.00,
			FotoKTP:      "foto_ktp_budi.jpg",
			FotoSelfie:   "foto_selfie_budi.jpg",
		},
		{
			ID:           "user-e17227f6-19d5-4403-a12b-be8a77c74897",
			NIK:          "b48b5b096d73aced5e94fa3829dcd8033c86b28b561ca4322c7fbe2193f6e5e5",
			FullName:     "Annisa",
			LegalName:    "Annisa Putri",
			TempatLahir:  "Bandung",
			TanggalLahir: customeparse.ParseTime("2001-02-02"),
			Gaji:         6700000.00,
			FotoKTP:      "foto_ktp_annisa.jpg",
			FotoSelfie:   "foto_selfie_annisa.jpg",
		},
	}
	for _, c := range consumers {
		result := db.Create(&c)
		if result.Error != nil {
			log.Println("Error seeding data:", result.Error)
			return
		}
	}
}

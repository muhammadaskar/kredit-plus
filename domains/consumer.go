package domains

import "time"

type Consumer struct {
	ID           string    `json:"id"`
	NIK          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Gaji         float64   `json:"gaji"`
	FotoKTP      string    `json:"foto_ktp"`
	FotoSelfie   string    `json:"foto_selfie"`
}

type GetConsumerID struct {
	ID string `uri:"id" binding:"required"`
}

package domains

import "time"

type Transaction struct {
	ID            string    `json:"id"`
	NomorKontrak  int       `json:"nomor_kontrak"`
	ConsumerID    string    `json:"consumer_id"`
	OTR           float64   `json:"otr"`
	AdminFee      float64   `json:"admin_fee"`
	JumlahCicilan int64     `json:"jumlah_cicilan"`
	JumlahBunga   float64   `json:"jumlah_bunga"`
	NamaAsset     string    `json:"nama_asset"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateTransaction struct {
	ConsumerID    string  `json:"consumer_id"`
	Tenor         int     `json:"tenor"`
	OTR           float64 `json:"otr"`
	AdminFee      float64 `json:"admin_fee"`
	JumlahCicilan int64   `json:"jumlah_cicilan"`
	JumlahBunga   float64 `json:"jumlah_bunga"`
	NamaAsset     string  `json:"nama_asset"`
}

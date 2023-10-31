package domains

type Tenor struct {
	ID          string `json:"id"`
	ConsumerID  string `json:"consumer_id"`
	LimitTenor1 int64  `json:"limit_tenor_1"`
	LimitTenor2 int64  `json:"limit_tenor_2"`
	LimitTenor3 int64  `json:"limit_tenor_3"`
	LimitTenor4 int64  `json:"limit_tenor_4"`
}

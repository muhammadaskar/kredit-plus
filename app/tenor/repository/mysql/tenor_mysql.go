package mysql

import (
	"github.com/muhammadaskar/kredit-plus/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindByConsumerId(consumerId string) (domains.Tenor, error)
	Update(consumerId string, tenorTo string, value int64) (domains.Tenor, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByConsumerId(consumerId string) (domains.Tenor, error) {
	var tenor domains.Tenor

	err := r.db.Where("consumer_id = ?", consumerId).Find(&tenor).Error
	if err != nil {
		return tenor, err
	}
	return tenor, nil
}

func (r *repository) Update(consumerId string, tenorTo string, value int64) (domains.Tenor, error) {
	var tenor domains.Tenor
	err := r.db.Model(&tenor).Where("consumer_id = ?", consumerId).Update("limit_tenor"+tenorTo, value).Error
	if err != nil {
		return tenor, err
	}
	return tenor, nil
}

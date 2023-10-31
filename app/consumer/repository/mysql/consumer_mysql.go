package mysql

import (
	"github.com/muhammadaskar/kredit-plus/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindById(id string) (domains.Consumer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindById(id string) (domains.Consumer, error) {
	var consumer domains.Consumer
	err := r.db.Where("id = ?", id).Find(&consumer).Error
	if err != nil {
		return consumer, err
	}
	return consumer, nil
}

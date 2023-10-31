package usecase

import (
	"github.com/muhammadaskar/kredit-plus/app/consumer/repository/mysql"
	"github.com/muhammadaskar/kredit-plus/domains"
	customsecurity "github.com/muhammadaskar/kredit-plus/utils/custom_security"
)

type ConsumerUseCase interface {
	FindById(consumer domains.GetConsumerID) (domains.Consumer, error)
}

type usecase struct {
	repository mysql.Repository
}

func NewUseCase(repository mysql.Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) FindById(consumerId domains.GetConsumerID) (domains.Consumer, error) {
	consumer, err := u.repository.FindById(consumerId.ID)
	if err != nil {
		return consumer, err
	}

	key := []byte("$3cReT!k3Y!60Hyh")

	decryptedNIK, err := customsecurity.Decrypt(key, consumer.NIK)
	if err != nil {
		return consumer, err
	}

	consumer.NIK = decryptedNIK
	return consumer, nil
}

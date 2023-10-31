package usecase

import (
	"github.com/muhammadaskar/kredit-plus/domains"
	"github.com/stretchr/testify/mock"
)

type MockTransactionUseCase struct {
	mock.Mock
}

func (m *MockTransactionUseCase) CreateTransaction(input domains.CreateTransaction) (domains.Transaction, error) {
	args := m.Called(input)
	return args.Get(0).(domains.Transaction), args.Error(1)
}

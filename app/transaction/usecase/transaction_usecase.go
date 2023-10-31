package usecase

import (
	"errors"
	"sync"

	"strconv"

	consumerMysql "github.com/muhammadaskar/kredit-plus/app/consumer/repository/mysql"
	tenorMysql "github.com/muhammadaskar/kredit-plus/app/tenor/repository/mysql"
	transactionMysql "github.com/muhammadaskar/kredit-plus/app/transaction/repository/mysql"
	"github.com/muhammadaskar/kredit-plus/domains"
	customgenerate "github.com/muhammadaskar/kredit-plus/utils/custom_generate"
)

type TransactionUseCase interface {
	CreateTransaction(input domains.CreateTransaction) (domains.Transaction, error)
}

type usecase struct {
	transactionRepo transactionMysql.Repository
	tenorRepo       tenorMysql.Repository
	consumerRepo    consumerMysql.Repository
}

func NewUseCase(transactionRepo transactionMysql.Repository, tenorRepo tenorMysql.Repository, consumerRepo consumerMysql.Repository) *usecase {
	return &usecase{transactionRepo, tenorRepo, consumerRepo}
}

func (u *usecase) CreateTransaction(input domains.CreateTransaction) (domains.Transaction, error) {
	transaction := domains.Transaction{}
	inputConsumerId := input.ConsumerID

	var wg sync.WaitGroup
	var consumer domains.Consumer
	var consumerError error
	var tenor domains.Tenor
	var tenorError error

	wg.Add(2)

	go func() {
		defer wg.Done()
		consumer, consumerError = u.consumerRepo.FindById(inputConsumerId)
	}()

	go func() {
		defer wg.Done()
		tenor, tenorError = u.tenorRepo.FindByConsumerId(inputConsumerId)
	}()

	wg.Wait()

	if consumerError != nil {
		return transaction, consumerError
	}

	if tenorError != nil {
		return transaction, tenorError
	}

	var tenorTo string
	var valueLimit int64
	switch input.Tenor {
	case 1:
		if input.JumlahCicilan > tenor.LimitTenor1 {
			return transaction, errors.New("EXCEEDS LIMIT FOR TENOR 1")
		}
		tenorTo = "1"
		valueLimit = tenor.LimitTenor1 - input.JumlahCicilan
	case 2:
		if input.JumlahCicilan > tenor.LimitTenor2 {
			return transaction, errors.New("EXCEEDS LIMIT FOR TENOR 2")
		}
		tenorTo = "2"
		valueLimit = tenor.LimitTenor2 - input.JumlahCicilan
	case 3:
		if input.JumlahCicilan > tenor.LimitTenor3 {
			return transaction, errors.New("EXCEEDS LIMIT FOR TENOR 3")
		}
		tenorTo = "3"
		valueLimit = tenor.LimitTenor3 - input.JumlahCicilan
	case 4:
		if input.JumlahCicilan > tenor.LimitTenor4 {
			return transaction, errors.New("EXCEEDS LIMIT FOR TENOR 4")
		}
		tenorTo = "4"
		valueLimit = tenor.LimitTenor4 - input.JumlahCicilan
	default:
		return transaction, errors.New("INVALID NUMBER")
	}

	_, err := u.tenorRepo.Update(consumer.ID, tenorTo, valueLimit)
	if err != nil {
		return transaction, err
	}

	transaction.ID = "transaction-" + customgenerate.GenerateTransactionCode()
	contractNumber, err := strconv.Atoi(customgenerate.GenerateContractNumber())
	if err != nil {
		return transaction, err
	}
	transaction.NomorKontrak = contractNumber
	transaction.ConsumerID = consumer.ID
	transaction.OTR = input.OTR
	transaction.AdminFee = input.AdminFee
	transaction.JumlahCicilan = input.JumlahCicilan
	transaction.JumlahBunga = input.JumlahBunga
	transaction.NamaAsset = input.NamaAsset

	newTransaction, err := u.transactionRepo.Create(transaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}

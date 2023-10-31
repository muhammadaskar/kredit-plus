package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/kredit-plus/app/transaction/usecase"
	"github.com/muhammadaskar/kredit-plus/domains"
	"github.com/stretchr/testify/assert"
)

func TestTransactionHandler_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(usecase.MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	router := gin.Default()
	router.POST("/api/v1/transactions", handler.Create)

	input := domains.CreateTransaction{
		ConsumerID:    "user-6af79619-c66f-4df7-98a2-b0ee282a3cc1",
		Tenor:         2,
		OTR:           1500000,
		AdminFee:      5000,
		JumlahCicilan: 200000,
		JumlahBunga:   10,
		NamaAsset:     "Samsung Galaxy A 01",
	}

	expectedTransaction := domains.Transaction{
		ID:            "transaction-KxT910hedD4J3LoL",
		NomorKontrak:  71679071,
		ConsumerID:    "user-6af79619-c66f-4df7-98a2-b0ee282a3cc1",
		OTR:           1500000,
		AdminFee:      5000,
		JumlahCicilan: 200000,
		JumlahBunga:   10,
		NamaAsset:     "Samsung Galaxy A 01",
		CreatedAt:     time.Now(),
	}

	mockUseCase.On("CreateTransaction", input).Return(expectedTransaction, nil)

	inputJSON, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/transactions", bytes.NewBuffer(inputJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)

	messageValue := response["meta"].(map[string]interface{})["message"].(string)

	assert.Equal(t, "Success to create transaction", messageValue)

	data := response["data"].(map[string]interface{})
	assert.Equal(t, expectedTransaction.ConsumerID, data["consumer_id"])
	assert.Equal(t, expectedTransaction.OTR, data["otr"])
	assert.Equal(t, expectedTransaction.AdminFee, data["admin_fee"])
	assert.Equal(t, expectedTransaction.JumlahBunga, data["jumlah_bunga"])
	assert.Equal(t, expectedTransaction.NamaAsset, data["nama_asset"])
}

func TestTransactionHandler_Create_Failure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(usecase.MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	router := gin.Default()
	router.POST("/api/v1/transactions", handler.Create)

	input := domains.CreateTransaction{
		ConsumerID:    "user-6af79619-c66f-4df7-98a2-b0ee282a3cc1",
		Tenor:         2,
		OTR:           1500000,
		AdminFee:      5000,
		JumlahCicilan: 200000,
		JumlahBunga:   10,
		NamaAsset:     "Samsung Galaxy A 01",
	}

	expectedError := errors.New("Test error message")
	mockUseCase.On("CreateTransaction", input).Return(domains.Transaction{}, expectedError)

	inputJSON, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/transactions", bytes.NewBuffer(inputJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)
	messageValue := response["meta"].(map[string]interface{})["message"].(string)

	assert.Equal(t, "Failed to create transaction", messageValue)
}

func TestTransactionHandler_Create_InsufficientTenor(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(usecase.MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	router := gin.Default()
	router.POST("/api/v1/transactions", handler.Create)

	input := domains.CreateTransaction{
		ConsumerID:    "user-6af79619-c66f-4df7-98a2-b0ee282a3cc1",
		Tenor:         1,
		OTR:           1500000,
		AdminFee:      5000,
		JumlahCicilan: 200000,
		JumlahBunga:   10,
		NamaAsset:     "Samsung Galaxy A 01",
	}

	mockUseCase.On("CreateTransaction", input).Return(domains.Transaction{}, errors.New("EXCEEDS LIMIT FOR TENOR 1"))

	inputJSON, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/transactions", bytes.NewBuffer(inputJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)
	data := response["data"].(map[string]interface{})
	errorsData := data["errors"].(string)
	assert.Equal(t, "EXCEEDS LIMIT FOR TENOR 1", errorsData)
}

func TestTransactionHandler_Create_InsufficientTenor2(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(usecase.MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	router := gin.Default()
	router.POST("/api/v1/transactions", handler.Create)

	input := domains.CreateTransaction{
		ConsumerID:    "user-6af79619-c66f-4df7-98a2-b0ee282a3cc1",
		Tenor:         2,
		OTR:           1500000,
		AdminFee:      5000,
		JumlahCicilan: 200000,
		JumlahBunga:   10,
		NamaAsset:     "Samsung Galaxy A 01",
	}

	mockUseCase.On("CreateTransaction", input).Return(domains.Transaction{}, errors.New("EXCEEDS LIMIT FOR TENOR 2"))

	inputJSON, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/transactions", bytes.NewBuffer(inputJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)
	data := response["data"].(map[string]interface{})
	errorsData := data["errors"].(string)
	assert.Equal(t, "EXCEEDS LIMIT FOR TENOR 2", errorsData)
}

func TestTransactionHandler_Create_InsufficientTenor3(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(usecase.MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	router := gin.Default()
	router.POST("/api/v1/transactions", handler.Create)

	input := domains.CreateTransaction{
		ConsumerID:    "user-6af79619-c66f-4df7-98a2-b0ee282a3cc1",
		Tenor:         3,
		OTR:           1500000,
		AdminFee:      5000,
		JumlahCicilan: 200000,
		JumlahBunga:   10,
		NamaAsset:     "Samsung Galaxy A 01",
	}

	mockUseCase.On("CreateTransaction", input).Return(domains.Transaction{}, errors.New("EXCEEDS LIMIT FOR TENOR 3"))

	inputJSON, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/transactions", bytes.NewBuffer(inputJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)
	data := response["data"].(map[string]interface{})
	errorsData := data["errors"].(string)
	assert.Equal(t, "EXCEEDS LIMIT FOR TENOR 3", errorsData)
}

func TestTransactionHandler_Create_InsufficientTenor4(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(usecase.MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	router := gin.Default()
	router.POST("/api/v1/transactions", handler.Create)

	input := domains.CreateTransaction{
		ConsumerID:    "user-6af79619-c66f-4df7-98a2-b0ee282a3cc1",
		Tenor:         4,
		OTR:           1500000,
		AdminFee:      5000,
		JumlahCicilan: 200000,
		JumlahBunga:   10,
		NamaAsset:     "Samsung Galaxy A 01",
	}

	mockUseCase.On("CreateTransaction", input).Return(domains.Transaction{}, errors.New("EXCEEDS LIMIT FOR TENOR 4"))

	inputJSON, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/transactions", bytes.NewBuffer(inputJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)
	data := response["data"].(map[string]interface{})
	errorsData := data["errors"].(string)
	assert.Equal(t, "EXCEEDS LIMIT FOR TENOR 4", errorsData)
}

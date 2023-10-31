package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/kredit-plus/app/transaction/usecase"
	"github.com/muhammadaskar/kredit-plus/domains"
	customresponse "github.com/muhammadaskar/kredit-plus/utils/custom_response"
)

type TransactionHandler struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionHandler(transactionUseCase usecase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{transactionUseCase}
}

func (h *TransactionHandler) Create(c *gin.Context) {
	var input domains.CreateTransaction

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTransaction, err := h.transactionUseCase.CreateTransaction(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to create transaction", http.StatusCreated, "success", newTransaction)
	c.JSON(http.StatusCreated, response)
}

package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/kredit-plus/app/consumer/usecase"
	"github.com/muhammadaskar/kredit-plus/domains"
	customresponse "github.com/muhammadaskar/kredit-plus/utils/custom_response"
)

type ConsumerHandler struct {
	consumerUseCase usecase.ConsumerUseCase
}

func NewConsumerHandler(consumerUsecase usecase.ConsumerUseCase) *ConsumerHandler {
	return &ConsumerHandler{consumerUsecase}
}

func (h *ConsumerHandler) FindById(c *gin.Context) {
	var input domains.GetConsumerID

	err := c.BindUri(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to get consumer", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	getConsumer, err := h.consumerUseCase.FindById(input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to get consumer", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := customresponse.APIResponse("Success to get consumer", http.StatusOK, "success", getConsumer)
	c.JSON(http.StatusOK, response)
}

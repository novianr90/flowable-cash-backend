package handlerUpdateTransaction

import (
	updateTransaction "flowable-cash-backend/controllers/transaction-controllers/update"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service updateTransaction.Service
}

func NewHandlerUpdateTransaction(service updateTransaction.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateTransaction(c *gin.Context) {
	var input updateTransaction.InputUpdateTransaction

	if err := c.ShouldBindQuery(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.UpdateTransactionService(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction sucessfully updated",
		"transaction": result,
	})
}

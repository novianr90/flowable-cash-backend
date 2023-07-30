package handlerDeleteTransaction

import (
	deleteTransaction "flowable-cash-backend/controllers/transaction-controllers/delete"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service deleteTransaction.Service
}

func NewHandlerDeleteTransaction(service deleteTransaction.Service) *handler {
	return &handler{service: service}
}

func (h *handler) DeleteTransaction(c *gin.Context) {
	var input deleteTransaction.InputDeleteTransaction

	if err := c.ShouldBindQuery(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.DeleteTransactionService(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction sucessfully deleted",
		"status":  "success",
	})
}

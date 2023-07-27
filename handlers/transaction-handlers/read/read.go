package handlerReadTransaction

import (
	readTransaction "flowable-cash-backend/controllers/transaction-controllers/read"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service readTransaction.Service
}

func NewHandlerReadTransaction(service readTransaction.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAllTransactions(c *gin.Context) {
	pengeluaran, pemasukkan, err := h.service.ReadAllTransactions()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pengeluaran": pengeluaran,
		"pemasukkan":  pemasukkan,
		"status":      "success",
	})
}

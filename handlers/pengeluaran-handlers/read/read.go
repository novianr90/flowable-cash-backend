package handlerReadPurchase

import (
	readPurchase "flowable-cash-backend/controllers/pengeluaran-controllers/read"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service readPurchase.Service
}

func NewReadPurchaseHandler(service readPurchase.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetPurchaseTransactions(c *gin.Context) {
	result, err := h.service.ReadAllPurchaseTypeTransactions()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pengeluaran": result,
	})
}

func (h *handler) GetPurchaseTransactionById(c *gin.Context) {
	var input readPurchase.InputReadPurchaseTransaction

	if err := c.ShouldBindQuery(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.ReadPurchaseTypeById(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pengeluaran": result,
	})
}

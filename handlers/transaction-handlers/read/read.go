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
	result, err := h.service.ReadAllTransactions()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": result,
	})
}

func (h *handler) GetTransactionById(c *gin.Context) {
	var input readTransaction.InputReadTransaction

	if err := c.ShouldBindQuery(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.ReadTransactionById(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction": result,
	})
}

func (h *handler) GetSaleTransactions(c *gin.Context) {
	result, err := h.service.ReadAllSaleTypeTransactions()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sale": result,
	})
}

func (h *handler) GetSaleTransactionById(c *gin.Context) {
	var input readTransaction.InputReadTransaction

	if err := c.ShouldBindQuery(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.ReadSaleTypeById(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sale": result,
	})
}

func (h *handler) GetPurchaseTransactions(c *gin.Context) {
	result, err := h.service.ReadAllPurchaseTypeTransactions()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"purchase": result,
	})
}

func (h *handler) GetPurchaseTransactionById(c *gin.Context) {
	var input readTransaction.InputReadTransaction

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
		"purchase": result,
	})
}

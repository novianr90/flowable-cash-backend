package handlerReadSaleTransaction

import (
	readSale "flowable-cash-backend/controllers/sale-controllers/read"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service readSale.Service
}

func NewReadSaleHandler(service readSale.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAllSaleTransactions(c *gin.Context) {
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
	var input readSale.InputReadSaleTransaction

	if err := c.ShouldBindQuery(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.ReadSaleTypeById(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sale": res,
	})
}

package handlerUpdateSaleTransaction

import (
	updateSale "flowable-cash-backend/controllers/sale-controllers/update"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service updateSale.Service
}

func NewUpdateSaleTransactionHandler(service updateSale.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateSaleTransactionById(c *gin.Context) {
	var input updateSale.InputUpdateSale

	if err := c.ShouldBindQuery(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.UpdateSaleTransaction(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction sucesfully updated",
		"sale":    res,
	})
}

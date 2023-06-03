package handlerUpdatePurchase

import (
	updatePurchase "flowable-cash-backend/controllers/purchase-controllers/update"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service updatePurchase.Service
}

func NewUpdatePurchaseHandler(service updatePurchase.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdatePurchaseTransactionById(c *gin.Context) {
	var input updatePurchase.InputUpdatePurchase

	if err := c.ShouldBindQuery(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.UpdatePurchaseTransactionById(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction successfully updated",
		"transaction": res,
	})
}

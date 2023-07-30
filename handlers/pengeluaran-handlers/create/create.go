package handlerCreatePurchase

import (
	createPurchase "flowable-cash-backend/controllers/pengeluaran-controllers/create"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service createPurchase.Service
}

func NewHandlerCreatePurchase(service createPurchase.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreatePemasukkan(ctx *gin.Context) {
	var input createPurchase.InputCreateTransaction

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.CreateTransactionService(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"pengeluaran": res,
		"status":      "success",
	})
}

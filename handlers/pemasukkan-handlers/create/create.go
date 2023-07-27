package handlerCreateSaleTransaction

import (
	createSale "flowable-cash-backend/controllers/pemasukkan-controllers/create"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service createSale.Service
}

func NewHandlerCreateSale(service createSale.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreatePemasukkan(ctx *gin.Context) {
	var input createSale.InputCreateTransaction

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
		"pemasukkan": res,
		"status":     "success",
	})
}

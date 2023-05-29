package handlerCreateTransaction

import (
	createTransaction "flowable-cash-backend/controllers/transaction-controllers/create"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service createTransaction.Service
}

func NewHandlerCreateTransaction(service createTransaction.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateTransaction(ctx *gin.Context) {
	var input createTransaction.InputCreateTransaction

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.CreateTransactionService(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"transaction": result,
	})
}

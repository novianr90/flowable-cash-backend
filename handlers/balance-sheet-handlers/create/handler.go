package handlerCreateBalanceSheet

import (
	createBalanceSheet "flowable-cash-backend/controllers/balance-sheet-controllers/create"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service createBalanceSheet.Service
}

func NewCreateBalanceSheetHandler(service createBalanceSheet.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateBalanceSheet(ctx *gin.Context) {
	var input createBalanceSheet.InputCreateBalanceSheet

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.CreateBalanceSheet(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"balance_sheet": res,
	})
}

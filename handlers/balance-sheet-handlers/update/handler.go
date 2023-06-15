package handlerUpdateBalanceSheet

import (
	updateBalanceSheet "flowable-cash-backend/controllers/balance-sheet-controllers/update"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service updateBalanceSheet.Service
}

func NewUpdateBalanceSheetHandler(service updateBalanceSheet.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateBalanceSheet(ctx *gin.Context) {
	var input updateBalanceSheet.InputUpdateBalanceSheet

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.UpdateBalanceSheet(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "Data succesfully updated",
		"balance_sheet": res,
	})
}

package handlerReadBalanceSheet

import (
	readBalanceSheet "flowable-cash-backend/controllers/balance-sheet-controllers/read"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service readBalanceSheet.Service
}

func NewReadBalanceSheetService(service readBalanceSheet.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetBalanceSheet(ctx *gin.Context) {
	result, err := h.service.GetBalanceSheet()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"balance_sheet": result,
	})
}

func (h *handler) GetBalanceSheetByAccountName(ctx *gin.Context) {
	var input readBalanceSheet.InputReadBalanceSheet

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.GetBalanceSheetByAccountName(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"balance_sheet": result,
	})
}

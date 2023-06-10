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

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

func (h *handler) UpdateAccount(ctx *gin.Context) {
	var input updateBalanceSheet.InputUpdateBalanceSheet

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	res, err := h.service.UpdateAccount(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"balance_sheet": res,
	})
}

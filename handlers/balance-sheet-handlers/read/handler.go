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

func (h *handler) GetAllAccounts(ctx *gin.Context) {

	var input *readBalanceSheet.InputReadBalanceSheet

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	result, err := h.service.GetAllAccounts(input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"balance_sheet": result,
	})
}

func (h *handler) GetAllAccountsByAccountName(ctx *gin.Context) {
	var input readBalanceSheet.InputReadBalanceSheet

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	result, err := h.service.GetAllAccountsByAccountName(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"balance_sheet": result,
	})
}

func (h *handler) GetAllSpecificAccounts(ctx *gin.Context) {
	var input readBalanceSheet.InputReadBalanceSheet

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	result, err := h.service.GetAllSpecificAccounts(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"balance_sheet": result,
	})
}

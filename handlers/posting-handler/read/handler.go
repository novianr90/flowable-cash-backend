package readPostingHandler

import (
	readPosting "flowable-cash-backend/controllers/posting-controller/read"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service readPosting.Service
}

func NewReadPostingHandler(service readPosting.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ReadAllPosting(ctx *gin.Context) {
	res, err := h.service.ReadAllPosting()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"posting": res,
	})
}

func (h *handler) ReadPostingByTrxIDAndAccountID(ctx *gin.Context) {

	var input readPosting.InputRead

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	res, err := h.service.ReadPostingByTrxIDAndAccountID(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"posting": res,
	})

}

package createPostingHandler

import (
	createPosting "flowable-cash-backend/controllers/posting-controller/create"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	createPosting createPosting.Service
}

func NewCreatePostingHandler(service createPosting.Service) *handler {
	return &handler{createPosting: service}
}

func (h *handler) RecordNewPosting(ctx *gin.Context) {
	var input createPosting.ModelCreatePosting

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	res, err := h.createPosting.CreateNewRecord(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"posting": res,
	})
}

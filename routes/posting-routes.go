package routes

import (
	createPosting "flowable-cash-backend/controllers/posting-controller/create"
	readPosting "flowable-cash-backend/controllers/posting-controller/read"

	createPostingHandler "flowable-cash-backend/handlers/posting-handler/create"
	readPostingHandler "flowable-cash-backend/handlers/posting-handler/read"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPostingRoutes(db *gorm.DB, routes *gin.RouterGroup) {

	// Create Dep
	createRepo := createPosting.NewCreatePostingRepo(db)
	createService := createPosting.NewCreatePostingService(createRepo)
	createHandler := createPostingHandler.NewCreatePostingHandler(createService)

	// Read Dep
	readRepo := readPosting.NewReadPostingRepository(db)
	readService := readPosting.NewReadPostingService(readRepo)
	readHandler := readPostingHandler.NewReadPostingHandler(readService)

	groupRouter := routes.Group("/posting")

	// Create
	groupRouter.POST("", createHandler.RecordNewPosting)

	// GET
	groupRouter.GET("", readHandler.ReadAllPosting)
	groupRouter.GET("/", readHandler.ReadPostingByTrxIDAndAccountID)
}

package summary

import (
	"portfolio-api/modules/summary/summary_handler"
	"portfolio-api/modules/summary/summary_repository"
	"portfolio-api/modules/summary/summary_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router registers the developer-scoped summary CMS endpoints.
func Router(developers *gin.RouterGroup, db *sqlx.DB) {
	repository := summary_repository.NewSummaryRepository(db)
	service := summary_service.NewSummaryService(repository)
	handler := summary_handler.NewSummaryHandler(service)

	developers.GET("/:userID/summary", handler.Get)
	developers.PUT("/:userID/summary", handler.Save)
}

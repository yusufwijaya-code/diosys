package language

import (
	"portfolio-api/modules/language/language_handler"
	"portfolio-api/modules/language/language_repository"
	"portfolio-api/modules/language/language_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router registers the developer-scoped language CMS endpoints.
func Router(developers *gin.RouterGroup, db *sqlx.DB) {
	repository := language_repository.NewLanguageRepository(db)
	service := language_service.NewLanguageService(repository)
	handler := language_handler.NewLanguageHandler(service)

	developers.GET("/:userID/languages", handler.GetByUser)
	developers.POST("/:userID/languages", handler.Create)
	developers.PUT("/:userID/languages/:id", handler.Update)
	developers.DELETE("/:userID/languages/:id", handler.Delete)
}

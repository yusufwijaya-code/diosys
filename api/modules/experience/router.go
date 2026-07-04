package experience

import (
	"portfolio-api/modules/experience/experience_handler"
	"portfolio-api/modules/experience/experience_repository"
	"portfolio-api/modules/experience/experience_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router registers the developer-scoped experience CMS endpoints.
func Router(developers *gin.RouterGroup, db *sqlx.DB) {
	repository := experience_repository.NewExperienceRepository(db)
	service := experience_service.NewExperienceService(repository)
	handler := experience_handler.NewExperienceHandler(service)

	developers.GET("/:userID/experiences", handler.GetByUser)
	developers.POST("/:userID/experiences", handler.Create)
	developers.PUT("/:userID/experiences/:id", handler.Update)
	developers.DELETE("/:userID/experiences/:id", handler.Delete)
}

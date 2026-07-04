package education

import (
	"portfolio-api/modules/education/education_handler"
	"portfolio-api/modules/education/education_repository"
	"portfolio-api/modules/education/education_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router registers the developer-scoped education CMS endpoints.
func Router(developers *gin.RouterGroup, db *sqlx.DB) {
	repository := education_repository.NewEducationRepository(db)
	service := education_service.NewEducationService(repository)
	handler := education_handler.NewEducationHandler(service)

	developers.GET("/:userID/educations", handler.GetByUser)
	developers.POST("/:userID/educations", handler.Create)
	developers.PUT("/:userID/educations/:id", handler.Update)
	developers.DELETE("/:userID/educations/:id", handler.Delete)
}

package skill

import (
	"portfolio-api/modules/skill/skill_handler"
	"portfolio-api/modules/skill/skill_repository"
	"portfolio-api/modules/skill/skill_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router registers the developer-scoped skill CMS endpoints.
func Router(developers *gin.RouterGroup, db *sqlx.DB) {
	repository := skill_repository.NewSkillRepository(db)
	service := skill_service.NewSkillService(repository)
	handler := skill_handler.NewSkillHandler(service)

	developers.GET("/:userID/skills", handler.GetByUser)
	developers.POST("/:userID/skills", handler.Create)
	developers.PUT("/:userID/skills/:id", handler.Update)
	developers.DELETE("/:userID/skills/:id", handler.Delete)
}

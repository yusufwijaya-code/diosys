package setting

import (
	"portfolio-api/modules/setting/setting_handler"
	"portfolio-api/modules/setting/setting_repository"
	"portfolio-api/modules/setting/setting_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router wires and registers the system setting endpoints.
func Router(public *gin.RouterGroup, protected *gin.RouterGroup, db *sqlx.DB) {
	repository := setting_repository.NewSettingRepository(db)
	service := setting_service.NewSettingService(repository)
	handler := setting_handler.NewSettingHandler(service)

	public.GET("/settings", handler.GetPublic)

	protectedGroup := protected.Group("/settings")
	{
		protectedGroup.GET("", handler.GetAll)
		protectedGroup.PUT("", handler.Update)
	}
}

package service

import (
	"portfolio-api/modules/service/service_handler"
	"portfolio-api/modules/service/service_repository"
	"portfolio-api/modules/service/service_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router wires and registers the agency service endpoints.
func Router(public *gin.RouterGroup, protected *gin.RouterGroup, db *sqlx.DB) {
	repository := service_repository.NewServiceRepository(db)
	svc := service_service.NewService(repository)
	handler := service_handler.NewServiceHandler(svc)

	public.GET("/services", handler.GetPublic)

	protectedGroup := protected.Group("/services")
	{
		protectedGroup.GET("", handler.GetAll)
		protectedGroup.GET("/:id", handler.GetByID)
		protectedGroup.POST("", handler.Create)
		protectedGroup.PUT("/:id", handler.Update)
		protectedGroup.DELETE("/:id", handler.Delete)
	}
}

package pricing

import (
	"portfolio-api/modules/pricing/pricing_handler"
	"portfolio-api/modules/pricing/pricing_repository"
	"portfolio-api/modules/pricing/pricing_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router wires and registers the pricing plan endpoints.
func Router(public *gin.RouterGroup, protected *gin.RouterGroup, db *sqlx.DB) {
	repository := pricing_repository.NewPricingRepository(db)
	service := pricing_service.NewPricingService(repository)
	handler := pricing_handler.NewPricingHandler(service)

	public.GET("/pricing", handler.GetPublic)

	protectedGroup := protected.Group("/pricing")
	{
		protectedGroup.GET("", handler.GetAll)
		protectedGroup.GET("/:id", handler.GetByID)
		protectedGroup.POST("", handler.Create)
		protectedGroup.PUT("/:id", handler.Update)
		protectedGroup.DELETE("/:id", handler.Delete)
	}
}

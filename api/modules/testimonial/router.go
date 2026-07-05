package testimonial

import (
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/modules/testimonial/testimonial_handler"
	"portfolio-api/modules/testimonial/testimonial_repository"
	"portfolio-api/modules/testimonial/testimonial_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Router(public *gin.RouterGroup, protected *gin.RouterGroup, db *sqlx.DB, gdrive *gdrive_helper.Client) {
	repo := testimonial_repository.NewTestimonialRepository(db)
	svc := testimonial_service.NewTestimonialService(repo, gdrive)
	handler := testimonial_handler.NewTestimonialHandler(svc)

	public.GET("/testimonials", handler.GetPublic)

	g := protected.Group("/testimonials")
	{
		g.GET("", handler.GetAll)
		g.GET("/:id", handler.GetByID)
		g.POST("", handler.Create)
		g.PUT("/:id", handler.Update)
		g.POST("/:id/photo", handler.UploadPhoto)
		g.DELETE("/:id", handler.Delete)
	}
}

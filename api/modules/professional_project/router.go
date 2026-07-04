package professional_project

import (
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/modules/professional_project/professional_project_handler"
	"portfolio-api/modules/professional_project/professional_project_repository"
	"portfolio-api/modules/professional_project/professional_project_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router registers public + CMS endpoints for professional projects.
func Router(public *gin.RouterGroup, cmsDevelopers *gin.RouterGroup, db *sqlx.DB, gdrive *gdrive_helper.Client) {
	repository := professional_project_repository.NewProfessionalProjectRepository(db)
	service := professional_project_service.NewProfessionalProjectService(repository, gdrive)
	handler := professional_project_handler.NewProfessionalProjectHandler(service)

	// Public detail endpoint
	public.GET("/professional-projects/:id", handler.GetPublicByID)

	// CMS developer-scoped endpoints
	cmsDevelopers.GET("/:userID/professional-projects", handler.GetByUser)
	cmsDevelopers.POST("/:userID/professional-projects", handler.Create)
	cmsDevelopers.DELETE("/:userID/professional-projects/:id", handler.Delete)
	cmsDevelopers.POST("/:userID/professional-projects/:id/thumbnail", handler.UploadThumbnail)
	cmsDevelopers.POST("/:userID/professional-projects/:id/features", handler.AddFeature)
	cmsDevelopers.DELETE("/:userID/professional-projects/:id/features/:featureID", handler.DeleteFeature)
	cmsDevelopers.POST("/:userID/professional-projects/:id/features/:featureID/images", handler.AddFeatureImage)
	cmsDevelopers.DELETE("/:userID/professional-projects/:id/features/:featureID/images/:imageID", handler.DeleteFeatureImage)
}

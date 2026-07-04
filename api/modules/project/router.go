package project

import (
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/modules/project/project_handler"
	"portfolio-api/modules/project/project_repository"
	"portfolio-api/modules/project/project_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router wires and registers the project endpoints.
func Router(public *gin.RouterGroup, developers *gin.RouterGroup, db *sqlx.DB, gdrive *gdrive_helper.Client) {
	repository := project_repository.NewProjectRepository(db)
	service := project_service.NewProjectService(repository, gdrive)
	handler := project_handler.NewProjectHandler(service)

	// Public company portfolio aggregated from all developers.
	public.GET("/projects", handler.GetAllPublic)
	public.GET("/projects/:id", handler.GetPublicByID)

	// Developer-scoped CMS management.
	developers.GET("/:userID/projects", handler.GetByUser)
	developers.POST("/:userID/projects", handler.Create)
	developers.PUT("/:userID/projects/:id", handler.Update)
	developers.DELETE("/:userID/projects/:id", handler.Delete)
	developers.POST("/:userID/projects/:id/thumbnail", handler.UploadThumbnail)
	developers.POST("/:userID/projects/:id/images", handler.AddImage)
	developers.DELETE("/:userID/projects/:id/images/:imageID", handler.DeleteImage)
}

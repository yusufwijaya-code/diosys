package certificate

import (
	"portfolio-api/modules/certificate/certificate_handler"
	"portfolio-api/modules/certificate/certificate_repository"
	"portfolio-api/modules/certificate/certificate_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router registers the developer-scoped certificate CMS endpoints.
func Router(developers *gin.RouterGroup, db *sqlx.DB) {
	repository := certificate_repository.NewCertificateRepository(db)
	service := certificate_service.NewCertificateService(repository)
	handler := certificate_handler.NewCertificateHandler(service)

	developers.GET("/:userID/certificates", handler.GetByUser)
	developers.POST("/:userID/certificates", handler.Create)
	developers.PUT("/:userID/certificates/:id", handler.Update)
	developers.DELETE("/:userID/certificates/:id", handler.Delete)
}

package developer

import (
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/modules/certificate/certificate_repository"
	"portfolio-api/modules/certificate/certificate_service"
	"portfolio-api/modules/developer/developer_handler"
	"portfolio-api/modules/developer/developer_service"
	"portfolio-api/modules/education/education_repository"
	"portfolio-api/modules/education/education_service"
	"portfolio-api/modules/experience/experience_repository"
	"portfolio-api/modules/experience/experience_service"
	"portfolio-api/modules/professional_project/professional_project_repository"
	"portfolio-api/modules/professional_project/professional_project_service"
	"portfolio-api/modules/project/project_repository"
	"portfolio-api/modules/project/project_service"
	"portfolio-api/modules/skill/skill_repository"
	"portfolio-api/modules/skill/skill_service"
	"portfolio-api/modules/summary/summary_repository"
	"portfolio-api/modules/summary/summary_service"
	"portfolio-api/modules/user/user_repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router wires the developer directory, public profile aggregate and CMS CRUD.
func Router(public *gin.RouterGroup, cmsDevelopers *gin.RouterGroup, db *sqlx.DB, gdrive *gdrive_helper.Client) {
	userRepository := user_repository.NewUserRepository(db)
	summaryService := summary_service.NewSummaryService(summary_repository.NewSummaryRepository(db))
	experienceService := experience_service.NewExperienceService(experience_repository.NewExperienceRepository(db))
	educationService := education_service.NewEducationService(education_repository.NewEducationRepository(db))
	certificateService := certificate_service.NewCertificateService(certificate_repository.NewCertificateRepository(db))
	skillService := skill_service.NewSkillService(skill_repository.NewSkillRepository(db))
	projectService := project_service.NewProjectService(project_repository.NewProjectRepository(db), gdrive)
	profProjService := professional_project_service.NewProfessionalProjectService(
		professional_project_repository.NewProfessionalProjectRepository(db), gdrive,
	)

	service := developer_service.NewDeveloperService(
		userRepository, summaryService, experienceService, educationService,
		certificateService, skillService, projectService, profProjService, gdrive,
	)
	handler := developer_handler.NewDeveloperHandler(service)

	// Public directory + profile aggregate.
	public.GET("/developers", handler.GetDirectory)
	public.GET("/developers/:username", handler.GetProfile)

	// CMS developer management.
	cmsDevelopers.GET("", handler.List)
	cmsDevelopers.POST("", handler.Create)
	cmsDevelopers.GET("/:userID", handler.GetByID)
	cmsDevelopers.PUT("/:userID", handler.Update)
	cmsDevelopers.DELETE("/:userID", handler.Delete)
	cmsDevelopers.POST("/:userID/photo", handler.UploadPhoto)
	cmsDevelopers.POST("/:userID/cv", handler.UploadCV)
}

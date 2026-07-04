package main

import (
	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/config"
	"portfolio-api/constants"
	"portfolio-api/middlewares"
	"portfolio-api/modules/auth"
	"portfolio-api/modules/certificate"
	"portfolio-api/modules/developer"
	"portfolio-api/modules/education"
	"portfolio-api/modules/experience"
	"portfolio-api/modules/message"
	"portfolio-api/modules/pricing"
	"portfolio-api/modules/project"
	"portfolio-api/modules/service"
	"portfolio-api/modules/setting"
	"portfolio-api/modules/skill"
	"portfolio-api/modules/summary"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// SetupRouter builds the gin engine and registers every module's routes.
func SetupRouter(
	db *sqlx.DB,
	cfg config.AppConfig,
	gdriveClient *gdrive_helper.Client,
	middleware middlewares.Middleware,
) *gin.Engine {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(middleware.PanicMiddleware(), gin.Logger(), middleware.CorsMiddleware())

	api := router.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Service is healthy", gin.H{"status": "up"})
	})

	// Authentication endpoints (Google sign-in).
	auth.Router(api, db, cfg, middleware)

	// Public (read-only) endpoints consumed by the Diosys website.
	public := api.Group("/public")

	// Protected CMS endpoints guarded by the JWT middleware.
	protected := api.Group("/cms", middleware.JwtMiddleware())

	// Developer-scoped CMS group: developer CRUD plus their nested portfolio.
	cmsDevelopers := protected.Group("/developers")

	developer.Router(public, cmsDevelopers, db, gdriveClient)
	summary.Router(cmsDevelopers, db)
	experience.Router(cmsDevelopers, db)
	education.Router(cmsDevelopers, db)
	certificate.Router(cmsDevelopers, db)
	skill.Router(cmsDevelopers, db)
	project.Router(public, cmsDevelopers, db, gdriveClient)

	// Agency-level endpoints.
	service.Router(public, protected, db)
	message.Router(public, protected, db)
	setting.Router(public, protected, db)
	pricing.Router(public, protected, db)

	return router
}

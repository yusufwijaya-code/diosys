package auth

import (
	"portfolio-api/config"
	"portfolio-api/middlewares"
	"portfolio-api/modules/auth/auth_handler"
	"portfolio-api/modules/auth/auth_service"
	"portfolio-api/modules/user/user_repository"
	"portfolio-api/modules/user/user_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router wires and registers the authentication endpoints.
func Router(apiGroup *gin.RouterGroup, db *sqlx.DB, cfg config.AppConfig, middleware middlewares.Middleware) {
	userRepository := user_repository.NewUserRepository(db)
	userService := user_service.NewUserService(userRepository)
	authService := auth_service.NewAuthService(userRepository, cfg)
	handler := auth_handler.NewAuthHandler(authService, userService, cfg)

	authGroup := apiGroup.Group("/auth")
	{
		authGroup.GET("/google/redirect", handler.GoogleRedirect)
		authGroup.GET("/google/callback", handler.GoogleCallback)
		authGroup.GET("/me", middleware.JwtMiddleware(), handler.Me)
	}
}

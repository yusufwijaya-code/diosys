package main

import (
	"fmt"
	"log"
	"os"

	"portfolio-api/base/helpers/gdrive_helper"
	"portfolio-api/config"
	"portfolio-api/database"
	"portfolio-api/middlewares"
	"portfolio-api/modules/user/user_repository"
	"portfolio-api/modules/user/user_service"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.AppConfig{}
	if err := config.LoadConfig(&cfg); err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	if err := os.Setenv("TZ", cfg.Timezone); err != nil {
		log.Printf("failed to set timezone: %v", err)
	}

	db := database.NewDB(cfg)
	InitMigration(db)

	seedWhitelistedAdmin(db, cfg)

	gdriveClient := initGoogleDrive(cfg)

	middleware := middlewares.NewMiddleware(cfg)
	router := SetupRouter(db, cfg, gdriveClient, middleware)

	address := cfg.AppHost + ":" + cfg.AppPort
	log.Printf("Server is starting on %s", address)
	if err := router.Run(address); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// seedWhitelistedAdmin ensures the single whitelisted administrator (who is also
// the first developer profile) exists.
func seedWhitelistedAdmin(db *sqlx.DB, cfg config.AppConfig) {
	repository := user_repository.NewUserRepository(db)
	service := user_service.NewUserService(repository)

	if err := service.EnsureWhitelistedAdmin(cfg.AdminEmail, "yusuf-wijaya", "Yusuf Wijaya"); err != nil {
		log.Printf("failed to seed whitelisted admin: %v", err)
	}
}

// initGoogleDrive builds the Google Drive client, returning nil when the
// credentials are incomplete so the API can still start without uploads.
func initGoogleDrive(cfg config.AppConfig) *gdrive_helper.Client {
	client, err := gdrive_helper.NewClient(gdrive_helper.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RefreshToken: cfg.GoogleDriveRefreshToken,
		RootFolderID: cfg.GoogleDriveFolderID,
	})
	if err != nil {
		log.Printf("google drive disabled: %v", err)
		return nil
	}
	log.Println("Google Drive client initialized")
	return client
}

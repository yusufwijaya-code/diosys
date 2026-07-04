package config

import "github.com/spf13/viper"

// AppConfig holds the application configuration loaded from the env file.
type AppConfig struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	AppHost        string `mapstructure:"APP_HOST"`
	AppPort        string `mapstructure:"APP_PORT"`
	AppFrontendUrl string `mapstructure:"APP_FRONTEND_URL"`
	Timezone       string `mapstructure:"TZ"`

	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`

	JwtSecret               string `mapstructure:"JWT_SECRET"`
	JwtTokenLifespanMinutes int    `mapstructure:"JWT_TOKEN_LIFESPAN_MINUTES"`

	// Google OAuth (Sign-in with Google, server-side redirect flow).
	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleRedirectURI  string `mapstructure:"GOOGLE_REDIRECT_URI"`

	// AdminEmail is the single whitelisted account allowed to sign in to the CMS.
	AdminEmail string `mapstructure:"ADMIN_EMAIL"`

	// Google Drive storage (shares the OAuth client credentials above).
	GoogleDriveRefreshToken string `mapstructure:"GOOGLE_DRIVE_REFRESH_TOKEN"`
	GoogleDriveFolderID     string `mapstructure:"GOOGLE_DRIVE_FOLDER_ID"`
}

// LoadConfig reads configuration from app.local.env, falling back to app.env.
func LoadConfig(cfg *AppConfig) error {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_PORT", "3005")
	viper.SetDefault("APP_FRONTEND_URL", "http://localhost:4200")
	viper.SetDefault("TZ", "Asia/Jakarta")
	viper.SetDefault("DB_HOST", "localhost:3306")
	viper.SetDefault("DB_NAME", "diosys_main")
	viper.SetDefault("DB_USERNAME", "root")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("JWT_TOKEN_LIFESPAN_MINUTES", 1440)
	viper.SetDefault("ADMIN_EMAIL", "yusufwijaya3@gmail.com")
	viper.SetDefault("GOOGLE_REDIRECT_URI", "http://localhost:3005/api/auth/google/callback")

	viper.SetConfigName("app.local")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigName("app")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	return viper.Unmarshal(cfg)
}

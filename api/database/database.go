package database

import (
	"fmt"
	"net/url"
	"time"

	"portfolio-api/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// NewDB opens a single MySQL connection pool for the portfolio database.
func NewDB(cfg config.AppConfig) *sqlx.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true&loc=%s&multiStatements=true",
		cfg.DbUsername,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbName,
		url.QueryEscape(cfg.Timezone),
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to open database %s: %v", cfg.DbName, err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping database %s: %v", cfg.DbName, err))
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	fmt.Printf("Database %s successfully connected\n", cfg.DbName)
	return db
}

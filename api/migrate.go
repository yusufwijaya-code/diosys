package main

import (
	"embed"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

// InitMigration applies all pending up migrations against the database.
func InitMigration(db *sqlx.DB) {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		log.Fatalf("failed to create migration driver: %v", err)
	}

	source, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		log.Fatalf("failed to load migration files: %v", err)
	}

	migrator, err := migrate.NewWithInstance("iofs", source, "mysql", driver)
	if err != nil {
		log.Fatalf("failed to initialize migrator: %v", err)
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("Database migrations applied successfully")
}

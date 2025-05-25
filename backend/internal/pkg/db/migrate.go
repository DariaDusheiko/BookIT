package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func RunMigrations(dsn string) error {
	m, err := migrate.New(
		"file://migrations",
		dsn)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	version, _, _ := m.Version()
	log.Printf("Database migrated to version: %d", version)

	return nil
}
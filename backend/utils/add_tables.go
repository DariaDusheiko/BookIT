package main

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/BookIT/backend/internal/app/models"
	"github.com/BookIT/backend/internal/app/repository"
)

func main() {
	// Initialize database connection
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate tables
	err = db.AutoMigrate(&models.Table{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	// Create repository instance
	tableRepo := repositories.NewTableRepository(db)

	// Define tables to insert
	tables := []models.Table{
		// Window tables
		{X: 15, Y: 15, Angle: 0, SeatsNumber: 2},
		{X: 15, Y: 35, Angle: 0, SeatsNumber: 2},
		{X: 15, Y: 55, Angle: 0, SeatsNumber: 4},
		{X: 15, Y: 75, Angle: 0, SeatsNumber: 4},

		// Standard tables
		{X: 40, Y: 25, Angle: 0, SeatsNumber: 4},
		{X: 40, Y: 50, Angle: 0, SeatsNumber: 6},
		{X: 40, Y: 75, Angle: 0, SeatsNumber: 4},

		// VIP tables
		{X: 65, Y: 20, Angle: 0, SeatsNumber: 8},
		{X: 65, Y: 50, Angle: 0, SeatsNumber: 6},
		{X: 65, Y: 80, Angle: 0, SeatsNumber: 8},

		// Bar tables
		{X: 90, Y: 15, Angle: 45, SeatsNumber: 2},
		{X: 90, Y: 30, Angle: 45, SeatsNumber: 2},
		{X: 90, Y: 45, Angle: 45, SeatsNumber: 2},

		// Corner tables
		{X: 85, Y: 70, Angle: -45, SeatsNumber: 4},
		{X: 85, Y: 90, Angle: -45, SeatsNumber: 4},
	}

	// Insert tables using repository
	ids, err := tableRepo.CreateTables(tables)
	if err != nil {
		log.Fatalf("Failed to insert tables: %v", err)
	}

	fmt.Printf("Successfully inserted %d tables with IDs: %v\n", len(ids), ids)
}
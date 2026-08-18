package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// PostgreSQL & GORM Database Initialization
func initDatabase(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		// Default PostgreSQL Connection DSN
		dsn = "host=localhost user=postgres password=postgres dbname=godb port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL database: %w", err)
	}

	// Auto-migrate GORM schemas to PostgreSQL
	err = db.AutoMigrate(&User{}, &Book{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto-migrate schemas: %w", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL Database & migrated tables!")
	DB = db
	return db, nil
}

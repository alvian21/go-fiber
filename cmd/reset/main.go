package main

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/seed"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.DatabaseInit()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	m, err := migrate.New(
		"file://database/migration/sql",
		dsn,
	)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	fmt.Println("Dropping all tables...")
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run down migrations: %v", err)
	}

	fmt.Println("Running up migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v", err)
	}

	fmt.Println("Running seeds...")
	seed.RunSeed()

	fmt.Println("Database reset and migration complete")
}

package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/database/seed"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()
	// RUN MIGRATION
	migration.RunMigration()
	// RUN SEED
	seed.RunSeed()

	app := fiber.New(fiber.Config{
		ReadBufferSize: 16384,
	})

	// INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":8080")
}

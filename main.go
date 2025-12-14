package main

import (
	"log"
	"path/filepath"

	"wekasir/config"
	"wekasir/database"
	"wekasir/entity"
	"wekasir/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// load env
	if err := godotenv.Load(filepath.Join(config.ProjectRootPath, ".env")); err != nil {
		log.Fatalf("failed load env: %s", err.Error())
	}

	// load database
	if err := database.DatabaseInit(); err != nil {
		log.Fatalf("failed connect db: %s", err.Error())
	}

	// migration
	if err := database.DB.AutoMigrate(
		&entity.User{},
	); err != nil {
		log.Fatalf("failed run migration: %s", err.Error())
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "*",
	}))
	app.Use(recover.New())

	routes.InitRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed run server: %v", err.Error())
	}
}

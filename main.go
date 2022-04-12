package main

import (
	"github.com/DiasOrazbaev/NEARFO/database"
	"github.com/DiasOrazbaev/NEARFO/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.DBConn = database.InitDB()
	app := fiber.New(fiber.Config{
		AppName: "NEARFood",
		Prefork: true,
	})
	app.Use(logger.New())
	app.Use(cors.New())
	routes.RoutesInit(app)
	app.Listen(":8080")
}

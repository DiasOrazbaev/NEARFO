package routes

import (
	"github.com/DiasOrazbaev/NEARFO/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "hello world",
		})
	})
	api := app.Group("/api")
	api.Get("/shops", controllers.GetShop)
	api.Post("/shops", controllers.CreateShop)
	api.Get("/shops/:id", controllers.GetShopById)
	api.Delete("/shops/:id", controllers.DeleteShop)
}

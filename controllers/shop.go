package controllers

import (
	"github.com/DiasOrazbaev/NEARFO/database"
	"github.com/DiasOrazbaev/NEARFO/models"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GetShop(ctx *fiber.Ctx) error {
	var shops []models.Shop
	database.DBConn.Find(&shops)
	return ctx.JSON(shops)
}

func GetShopById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var shop models.Shop
	database.DBConn.First(&shop, id)
	return ctx.JSON(shop)
}

func CreateShop(ctx *fiber.Ctx) error {
	var shop models.Shop
	err := ctx.BodyParser(&shop)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	id, _ := gonanoid.Generate("0123456789", 10)
	shop.ID = id
	database.DBConn.Create(&shop)
	return ctx.Status(200).JSON(shop)
}

func DeleteShop(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var shop models.Shop
	database.DBConn.Delete(&shop, id)
	return ctx.Status(200).SendString("Post deleted successfully")
}

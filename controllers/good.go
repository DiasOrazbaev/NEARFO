package controllers

import (
	"github.com/DiasOrazbaev/NEARFO/database"
	"github.com/DiasOrazbaev/NEARFO/models"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GetGood(ctx *fiber.Ctx) error {
	var goods []models.Good
	database.DBConn.Find(&goods)
	return ctx.JSON(goods)
}

func GetGoodById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var good models.Good
	database.DBConn.First(&good, id)
	return ctx.JSON(good)
}

func CreateGood(ctx *fiber.Ctx) error {
	var good models.Good
	err := ctx.BodyParser(&good)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	id, _ := gonanoid.Generate("0123456789", 10)
	good.ID = id
	database.DBConn.Create(&good)
	return ctx.Status(201).JSON(good)
}

func DeleteGood(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var good models.Good
	database.DBConn.Delete(&good, id)
	return ctx.Status(200).SendString("Post deleted successfully")
}

package services

import (
	"FIVERR/database"
	"FIVERR/models"

	"github.com/gofiber/fiber"
)

func GetCategory(c *fiber.Ctx) {
	db := database.DBConn
	var category []models.Category
	db.Find(&category)
	c.JSON(category)
}

func GetCategorys(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("CategoryId")

	var category []models.Category
	db.Find(&category, id)
	c.JSON(category)
}

func NewCategory(c *fiber.Ctx) {
	db := database.DBConn
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&category)
	c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) {
	category := new(models.Category)
	id := c.Params("id")
	db := database.DBConn
	if err := c.BodyParser(category); err != nil {
		c.Status(503).SendString(err.Error())
	}
	db.Where("CategoryId = ?", id).Updates(&category)
	c.Status(200).JSON(category)
}

func DeleteCategory(c *fiber.Ctx) {
	id := c.Params("CategoryId")
	db := database.DBConn

	var category models.Category
	db.First(&category, id)
	if category.Libeller == "" {
		c.Status(500).Send("No Category Found with ID")
		return
	}
	db.Delete(&category)
	c.Send("Category Successfully deleted")
}

package handlers

import (
	"api/models"

	"github.com/gofiber/fiber/v2"
)

var items = make(map[string]models.Item)

func GetItems(c *fiber.Ctx) error {
	itemList := make([]models.Item, 0, len(items))

	for _, item := range items {
		itemList = append(itemList, item)
	}

	return c.JSON(itemList)
}

func GetItem(c *fiber.Ctx) error {

	id := c.Params("id")
	item, exists := items[id]

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Item not found",
		})
	}

	return c.JSON(item)
}

func CreateItem(c *fiber.Ctx) error {
	var item models.Item

	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if item.ID == "" || item.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID and Name are required",
		})
	}

	if _, exists := items[item.ID]; exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Item already exists",
		})
	}

	items[item.ID] = item

	return c.Status(fiber.StatusCreated).JSON(item)
}

func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, exists := items[id]; !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Item not found",
		})
	}

	var item models.Item

	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	item.ID = id

	items[id] = item

	return c.JSON(item)
}

func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, exists := items[id]; !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Item not found",
		})
	}

	delete(items, id)

	return c.SendStatus(fiber.StatusNoContent)

}

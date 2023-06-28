package book

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	if err := c.SendString("All books"); err != nil {
		c.SendString(err.Error())
	}

	return nil
}

func GetBook(c *fiber.Ctx) error {
	if err := c.SendString("Single Book"); err != nil {
		c.SendString(err.Error())
	}

	return nil
}

func NewBook(c *fiber.Ctx) error {
	if err := c.SendString("New book"); err != nil {
		c.SendString(err.Error())
	}

	return nil
}

func DeleteBook(c *fiber.Ctx) error {
	if err := c.SendString("delete book"); err != nil {
		c.SendString(err.Error())
	}

	return nil
}

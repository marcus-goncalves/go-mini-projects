package book

import (
	"github.com/gofiber/fiber/v2"
	"projects.com/prj-fiber/internals/database"
	"projects.com/prj-fiber/internals/entities"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DbConn
	books := []entities.Book{}

	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	db := database.DbConn
	id := c.Params("id")
	book := entities.Book{}

	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DbConn
	book := entities.Book{
		Title:  "1984",
		Author: "George Orwell",
		Rating: 5,
	}

	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	db := database.DbConn
	id := c.Params("id")
	book := entities.Book{}

	db.First(&book, id)
	if book.Title == "" {
		c.Status(404).SendString("Book id not found")
		return nil
	}

	db.Delete(&book)
	return c.SendString("deleted")
}

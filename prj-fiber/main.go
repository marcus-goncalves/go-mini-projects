package main

import (
	"github.com/gofiber/fiber/v2"
	"projects.com/prj-fiber/internals/book"
)

const (
	url = "/api/v1"
)

func main() {
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get(url+"/book", book.GetBooks)
	app.Get(url+"/book/:id", book.GetBook)
	app.Post(url+"/book", book.NewBook)
	app.Delete(url+"/book/:id", book.DeleteBook)
}

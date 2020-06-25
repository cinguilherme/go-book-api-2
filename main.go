package main

import (
	"github.com/cinguilherme/go-api-fiber/database"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	database.GetModel()
	c.Send("hello world")
}

func setupBookRoutes(app *fiber.App) {
	app.Get("/api/v1/book", database.AllBooks)
	app.Post("/api/v1/book", database.CreateBook)
	app.Get("/api/v1/book/:id", database.GetBook)
	app.Delete("/api/v1/book/:id", database.DeleteBook)
}

func main() {
	app := fiber.New()
	database.PerformMigrations()
	app.Get("/", helloWorld)

	setupBookRoutes(app)

	app.Listen(3000)
}

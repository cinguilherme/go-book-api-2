package main

import (
	"fmt"

	"github.com/cinguilherme/go-api-fiber/book"
	"github.com/cinguilherme/go-api-fiber/database"
	"github.com/cinguilherme/go-api-fiber/video"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("hello world")
}

func setupBookRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.AllBooks)
	app.Post("/api/v1/book", book.CreateBook)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	database.DBConn = database.GetModel()
	// defer database.DBConn.Close()
	database.DBConn.AutoMigrate(&book.Book{})
	database.DBConn.AutoMigrate(&video.Video{})
	fmt.Println("DB migrated")

}

func main() {
	app := fiber.New()
	initDatabase()

	app.Get("/", helloWorld)

	setupBookRoutes(app)

	app.Listen(3000)
}

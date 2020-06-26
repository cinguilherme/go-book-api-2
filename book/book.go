package book

import (
	"fmt"

	"github.com/cinguilherme/go-api-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
	ID     int    `json:"id"`
}

func AllBooks(c *fiber.Ctx) {
	db := database.DBConn
	var all []Book
	db.Find(&all)
	c.JSON(all)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	fmt.Println(id)
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func CreateBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book
	book.Title = "LOrds"
	book.Author = "Tolkien"

	db.Create(&book).Commit()
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("delete books endpoint")
}

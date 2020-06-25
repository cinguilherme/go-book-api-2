package database

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Book struct {
	gorm.Model
	Metadata postgres.Jsonb
	Body     string
	ID       int
}

func PerformMigrations() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=bookuser dbname=books_db password=bookdbpass sslmode=disable")

	defer db.Close()
	if err != nil {
		panic("error connecting to db")
	}
	fmt.Println("db connected")
	db.AutoMigrate(&Book{})
}

func GetModel() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=bookuser dbname=books_db password=bookdbpass sslmode=disable")

	defer db.Close()
	if err != nil {
		panic("error connecting to db")
	}
	fmt.Println("db connected")

}

func AllBooks(c *fiber.Ctx) {
	c.Send("all books endpoint")
}

func GetBook(c *fiber.Ctx) {
	c.Send("single books endpoint")
}

func CreateBook(c *fiber.Ctx) {
	c.Send("create books endpoint")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("delete books endpoint")
}

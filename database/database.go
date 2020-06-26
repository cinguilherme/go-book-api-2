package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DBConn *gorm.DB
)

func GetModel() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=bookuser dbname=books_db password=bookdbpass sslmode=disable")

	// defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic("error connecting to db")
	}

	fmt.Println("db connected")
	db.LogMode(true)
	return db
}

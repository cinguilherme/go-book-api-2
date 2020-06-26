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
	connectionStr := "host=localhost port=5432 user=bookuser dbname=books_db password=bookdbpass sslmode=disable"
	// connectionStr := "host=book_db_go port=5432 user=bookuser dbname=books_db password=bookdbpass sslmode=disable"

	db, err := gorm.Open("postgres", connectionStr)

	// defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic("error connecting to db " + err.Error())
	}

	fmt.Println("db connected")
	db.LogMode(true)
	return db
}

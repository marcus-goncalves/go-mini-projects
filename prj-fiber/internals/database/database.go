package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"projects.com/prj-fiber/internals/entities"
)

var (
	DbConn *gorm.DB
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open("./internals/database/book.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&entities.Book{})
	fmt.Println("db migrated")

	DbConn = db
}

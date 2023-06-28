package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	DbConn *gorm.DB
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open("./internals/database/book.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DbConn = db
}

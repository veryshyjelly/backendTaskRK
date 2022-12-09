package database

import (
	"fmt"

	"github.com/veryshyvelly/task2-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBInstance(dbName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Student{})
	db.AutoMigrate(&models.Admin{})

	fmt.Println("Connection Opened to Database")

	return db
}

var DB *gorm.DB = DBInstance("test.db")

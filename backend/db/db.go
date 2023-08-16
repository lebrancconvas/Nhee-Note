package db

import (
	"fmt"
	"lebrancconvas/nheenote/forms"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", "localhost", "postgres", "P@ssw0rd", "nheenote", "2003")

	DBConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected!")

	DBConn.AutoMigrate(&forms.Dept{})
	fmt.Println("Migrated Database!")

}

func CloseDB() {
	db, err := DBConn.DB()
	if err != nil {
		panic(err)
	}
	db.Close()
	fmt.Println("Database Closed!")
}

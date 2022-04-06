package config

import (
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	postgresDB *gorm.DB
)

func getDBUrl() string {
	defaultDb := "postgres://book_store_backend:123456789@localhost:5432/book_store_db"
	var dbURL = os.Getenv("database")
	if dbURL != "" {
		defaultDb = dbURL
	}
	return defaultDb
}

func Connect() {
	db, err := gorm.Open(postgres.Open(getDBUrl()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	postgresDB = db
}

func GetDB() *gorm.DB {
	return postgresDB
}

package models

import (
	"BookStoreAPIs/core/config"
	"gorm.io/gorm"
	"log"
)

var dataBase *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	dataBase = config.GetDB()
	if err := dataBase.AutoMigrate(&Book{}); err != nil {
		panic(err)
	}
}

func (b *Book) CreateBook() *Book {
	if err := dataBase.Create(&b); err != nil {
		log.Fatal(err)
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	dataBase.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	dataBase := dataBase.Where("ID=?", id).Find(&book)
	return &book, dataBase
}

func DeleteBook(id int64) Book {
	var book Book
	dataBase.Where("ID=?", id).Delete(book)
	return book
}

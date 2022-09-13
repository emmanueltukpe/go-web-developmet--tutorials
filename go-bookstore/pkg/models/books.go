package models

import (
	"log"

	"github.com/emmanueltukpe/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Publication string
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatalf("an err occured setting up auto migrations: %v", err)
	}
}

func CreateBook(book *Book) Book {
	b := Book{
		Author:      book.Author,
		Name:        book.Name,
		Publication: book.Publication,
	}
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

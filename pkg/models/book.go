package models

import (
	"github.com/akash/bookApi/pkg/config"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string `gorm:""json:"name"`
	Author string `json:"author"`
	City   string `json:"city"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(bookId int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", bookId).Find(&book)
	return &book, db
}

func DeleteBook(bookId int64) Book {
	var book Book
	db.Where("ID=?", bookId).Delete(&book)
	return book
}
 
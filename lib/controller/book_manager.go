package controller

import (
	"go_practice/lib/models"
)

type BookManager struct {
	AllBooks []models.Book
}

func (bm *BookManager) AddBook(book *models.Book) {
   bm.AllBooks = append(bm.AllBooks, *book)
}
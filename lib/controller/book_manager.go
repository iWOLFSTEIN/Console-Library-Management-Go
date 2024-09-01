package controller

import (
	"go_practice/lib/models"
)

func contains(books *[]models.Book, title string) bool {
	for _, item := range *books {
		if item.Title == title {
			return true
		}
	}
	return false
}

type BookManager struct {
	AllBooks []models.Book
}

func (bm *BookManager) AddBook(book *models.Book) bool {
	bookToBeAdded := *book
	if contains(&bm.AllBooks, bookToBeAdded.Title) {
		return false
	}

	bm.AllBooks = append(bm.AllBooks, *book)
	return true
}

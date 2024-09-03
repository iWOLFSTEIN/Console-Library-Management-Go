package controller

import (
	"go_practice/lib/models"
)

func contains[T comparable](items *[]T, item T) bool {
	for _, element := range *items {
		if element == item {
			return true
		}
	}
	return false
}

type BookManager struct {
	AllUsers []models.User
	AllBooks []models.Book

	Allocations []models.Allocation
}

func (bm *BookManager) GetBookOrNil(id int) *models.Book {
	for _, book := range bm.AllBooks {
		if book.Id == id {
			return &book
		}
	}

	return nil
}

func (bm *BookManager) GetUserOrNil(id int) *models.User {
	for _, user := range bm.AllUsers {
		if user.Id == id {
			return &user
		}
	}

	return nil
}

func (bm *BookManager) BookExist(book *models.Book) bool {
	return contains(&bm.AllBooks, *book)
}

func (bm *BookManager) UserExist(user *models.User) bool {
	return contains(&bm.AllUsers, *user)
}

func (bm *BookManager) AddBook(book *models.Book) bool {
	if bm.BookExist(book) {
		return false
	}

	bm.AllBooks = append(bm.AllBooks, *book)
	return true
}

func (bm *BookManager) AddUser(user *models.User) bool {
	if bm.UserExist(user) {
		return false
	}

	bm.AllUsers = append(bm.AllUsers, *user)
	return true
}

func (bm *BookManager) BookAlreadyAllocated(book *models.Book, user *models.User) bool {
	isAlreadyAllocated := false
	for _, allocation := range bm.Allocations {
		if allocation.User == *user && allocation.Book == *book {
			isAlreadyAllocated = true
		}
	}
	return isAlreadyAllocated
}

func (bm *BookManager) AllocateBook(book *models.Book, user *models.User) string {
	if bm.BookAlreadyAllocated(book, user) {
		return "Book is already allocated to this user"
	}

	bm.Allocations = append(bm.Allocations, models.Allocation{User: *user, Book: *book})

	return "Book allocated"
}

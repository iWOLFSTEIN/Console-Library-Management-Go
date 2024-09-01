package main

import (
	"fmt"
	"go_practice/lib/controller"
	"go_practice/lib/models"
	"strconv"
)

func main() {
	fmt.Println("WELCOME!")
	bookManager := controller.BookManager{}
	optionsMap := map[int]string{1: "Register a book", 2: "Allocate a book", 3: "Search a book", 4: "Print all books"}
	for {
		for key, value := range optionsMap {
			fmt.Println(key, value)
		}
		var userInput string
		fmt.Scanln(&userInput)
		selectedIndex, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		switch selectedIndex {
		case 1:
			RegisterABook(&bookManager)
		case 4:
			PrintAllBooks(&bookManager)
		}
	}

}

func RegisterABook(bookManager *controller.BookManager) {
	var title string
	var author string
	var pages int

	println("Write title of the book")
	fmt.Scanln(&title)
	println("Write author name of this book")
	fmt.Scanln(&author)
	println("Number of pages this book has")
	var pagesStr string
	fmt.Scanln(&pagesStr)
	pages, err := strconv.Atoi(pagesStr)
	if err != nil {
		fmt.Println("Invalid page number")
		return
	}

	book := models.Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}

	isAdded := bookManager.AddBook(&book)
	if !isAdded {
		println(book.Title, "already exists!")
		return
	}

	fmt.Println(book.Title + " is added to the collection")

}

func PrintAllBooks(bookManager *controller.BookManager) {
	fmt.Printf("All Books: %+v\n", bookManager.AllBooks)
}

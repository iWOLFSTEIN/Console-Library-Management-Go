package main

import (
	"fmt"
	"go_practice/lib/controller"
	"go_practice/lib/models"
)

func main() {
	fmt.Println("hello world")

	bookManager := controller.BookManager{}

	book := models.Book{
		Title:  "A brief history of time",
		Author: "Stephen Hawking",
		Pages:  100,
	}

	bookManager.AddBook(&book)

	fmt.Println(book.Author)

	fmt.Printf("All Books: %+v\n", bookManager.AllBooks)
}

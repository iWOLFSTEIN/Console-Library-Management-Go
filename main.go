package main

import (
	"fmt"
	"go_practice/lib/controller"
	"go_practice/lib/models"
	"strconv"
)

func main() {
	userIdCounter := 0
	bookIdCounter := 0
	fmt.Println("WELCOME!")
	bookManager := controller.BookManager{}
	optionsMap := map[int]string{1: "Register a book", 2: "Allocate a book", 3: "Search a book", 4: "Print all books", 5: "Register a user"}
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
			RegisterABook(&bookManager, &bookIdCounter)
		case 2:
			AllocatedABook(&bookManager)
		case 3:
			SearchBooks(&bookManager)
		case 4:
			PrintAllBooks(&bookManager)
		case 5:
			RegisterAUser(&bookManager, &userIdCounter)
		}
	}

}

func SearchBooks(bookManager *controller.BookManager) {
	println("Write title of the book")
	var bookTitle string
	fmt.Scanln(&bookTitle)

	var allSearchedBooks *[]models.Book = bookManager.GetAllSearchedBooksResults(bookTitle)

	for _, book := range *allSearchedBooks {
		println(book.Id, book.Title)
	}

}

func AllocatedABook(bookManager *controller.BookManager) {
	println("Write id of the book")
	var bookIdStr string
	fmt.Scanln(&bookIdStr)
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		println("Invalid book id")
		return
	}
	book := bookManager.GetBookOrNil(bookId)
	if book == nil {
		println("Book with this id does not exist")
		return
	}
	println(book.Title, "is fetched")

	println("Write id of the user")
	var userIdStr string
	fmt.Scanln(&userIdStr)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		println("Invalid user id")
		return
	}
	user := bookManager.GetUserOrNil(userId)
	if user == nil {
		println("User with this id does not exist")
		return
	}
	println(user.Name, "is fetched")

	bookManager.AllocateBook(book, user)
	println("Book id", book.Id, "is assigned to user with id", user.Id)

}

func RegisterAUser(bookManager *controller.BookManager, userIdCounter *int) {
	var name string

	println("Write name of the user")
	fmt.Scanln(&name)

	user := models.User{Id: *userIdCounter, Name: name}

	isAdded := bookManager.AddUser(&user)
	if !isAdded {
		println(user.Name, "already exists!")
		return
	}

	fmt.Println(user.Name + " is registered")
	*userIdCounter += 1
}

func RegisterABook(bookManager *controller.BookManager, bookIdCounter *int) {
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
		Id:     *bookIdCounter,
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
	*bookIdCounter += 1

}

func PrintAllBooks(bookManager *controller.BookManager) {
	fmt.Printf("All Books: %+v\n", bookManager.AllBooks)
}

package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	model "github.com/Abzaek/library_management/models"
	"github.com/Abzaek/library_management/services"
)

func newLineRemove(str string) string {
	var builder strings.Builder

	for _, s := range str {
		if s != '\n' {
			builder.WriteRune(s)
		}
	}

	return builder.String()
}
func importBook() model.Book {
	var ID int
	var Title string
	var Author string
	var Status string
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("enter the book Title: ")
	Title, err1 := reader.ReadString('\n')

	fmt.Printf("enter the book ID: ")
	StringID, err2 := reader.ReadString('\n')

	StringID = newLineRemove(StringID)

	ID, errr := strconv.Atoi(StringID)
	if errr != nil {
		fmt.Printf("error: %s\n", errr)
	}

	fmt.Printf("enter the book's Author's name: ")
	Author, err3 := reader.ReadString('\n')

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Printf("error while reading \n")
	}

	Status = "Available"

	return model.Book{ID: ID, Title: Title, Author: Author, Status: Status}

}

func ControlStarter() {
	var lib services.Library = services.Library{AllBooks: make(map[int]model.Book), Members: make(map[int]model.Member)}
	cond := true

	for cond {
		fmt.Println(`
		Hello welcome! the following are the available commands:
			/Add		  				to add a book
			/Remove bookID 				to remove a book with ID "bookID"
			/Borrow bookID memberID  	to borrow a book with ID "bookID" and a member ID "memberID"
			/Return bookID memberID  	to borrow a book with ID "bookID" and a member ID "memberID"
			/List 	 					to list all the available books
			/ListBorrowed memberID  	to list all the borrowed books by a member with a member ID "memberID"
			/exit						to exit 
		`)

		reader := bufio.NewReader(os.Stdin)
		var inpt, err = reader.ReadString('\n')
		input := strings.Fields(inpt)

		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}

		switch input[0] {
		case "/Add":
			lib.AddBook(importBook())
		case "/Remove":
			bookID := StringParser1(input)
			lib.RemoveBook(bookID)
		case "/Borrow":
			bookID, memberID := StringParser2(input)
			lib.BorrowBook(bookID, memberID)
		case "/Return":
			bookID, memberID := StringParser2(input)
			lib.ReturnBook(bookID, memberID)
		case "/List":
			lib.ListAvailableBooks()
		case "/ListBorrowed":
			memberID := StringParser1(input)
			lib.ListBorrowedBooks(memberID)
		case "/exit":
			cond = false
		default:
			fmt.Println("command unrecognised")
		}

		var inpt2 string
		fmt.Printf("wanna continue? (y/n): ")
		fmt.Scan(&inpt2)

		if !(strings.HasPrefix(inpt2, "y") || strings.HasPrefix(inpt2, "Y")) {
			// fmt.Println(strings.HasPrefix(inpt2, "y"))
			cond = false
		}

	}

}

func StringParser1(command []string) int {
	intValue, err := strconv.Atoi(command[1])

	if err != nil {
		fmt.Printf("error: %s", err)
		return 0
	}
	return intValue
}

func StringParser2(command []string) (int, int) {
	bookID, err1 := strconv.Atoi(command[1])
	memberID, err2 := strconv.Atoi(command[2])
	if err1 != nil {
		fmt.Printf("error: %s", err1)
		return 0, 0
	} else if err2 != nil {
		fmt.Printf("error: %s", err2)
		return 0, 0
	}
	return bookID, memberID
}

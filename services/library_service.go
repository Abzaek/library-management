package services

import (
	"errors"
	"fmt"

	model "github.com/Abzaek/library_management/models"
)

type LibraryManager interface {
	AddBook(book model.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []model.Book
	ListBorrowedBooks(memberID int) []model.Book
}

type Library struct {
	AllBooks map[int]model.Book
	Members  map[int]model.Member
}

func (lib *Library) AddBook(book model.Book) {
	b, exists := lib.AllBooks[book.ID]

	if exists {
		fmt.Printf("the book %s already exists\n", b.Title)
		return
	}
	fmt.Println("Successfully Added")
	lib.AllBooks[book.ID] = book
}

func (lib *Library) RemoveBook(bookID int) {
	b, exists := lib.AllBooks[bookID]

	if !exists {
		fmt.Printf("The book with ID %d doesn't exist %s\n", bookID, b.Title)
		return
	}

	delete(lib.AllBooks, bookID)
	fmt.Println("Successfully Removed!")
}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	book := lib.AllBooks[bookID]

	if book.Status == "Not Available" {
		return errors.New(`book is not available`)
	}
	member := lib.Members[memberID]

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	book.Status = "Not Available"

	return errors.New("Successful")

}

func (lib *Library) ReturnBook(bookID int, memberId int) error {
	book := lib.AllBooks[bookID]
	book.Status = "Available"

	member := lib.Members[memberId]

	var newSlice []model.Book

	for _, b := range member.BorrowedBooks {
		if b.ID != bookID {
			newSlice = append(newSlice, b)
		}
	}

	member.BorrowedBooks = newSlice
	lib.Members[memberId] = member
	lib.AllBooks[bookID] = book
	return errors.New("successfully returned!")
}

func (lib *Library) ListAvailableBooks() []model.Book {
	newSlice := []model.Book{}

	for _, book := range lib.AllBooks {
		if book.Status == "Available" {
			newSlice = append(newSlice, book)
		}
	}
	return newSlice
}

func (lib *Library) ListBorrowedBooks(memberID int) []model.Book {
	return lib.Members[memberID].BorrowedBooks
}

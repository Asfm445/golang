package services

import (
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) map[int]models.Book
	Register(memberName string) int
	Login(memberID int) bool
}

type Library struct {
	store_book map[int]models.Book
	members    map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		store_book: make(map[int]models.Book),
		members:    make(map[int]models.Member),
	}
}

// Add the book
func (library *Library) AddBook(book models.Book) {
	book.ID = len(library.store_book) + 1
	book.Status = "available"
	library.store_book[book.ID] = book
}

// Delete the book
func (library *Library) RemoveBook(bookID int) {
	book, exists := library.store_book[bookID]
	if !exists {
		fmt.Printf("Book ID %d does not exist.\n", bookID)
		return
	}
	if book.Status == "borrowed" {
		fmt.Printf("Book ID %d is currently borrowed and cannot be removed.\n", bookID)
		return
	}
	delete(library.store_book, bookID)
	fmt.Printf("Book ID %d removed successfully.\n", bookID)
}

// Borrow the book
func (library *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExists := library.store_book[bookID]
	if !bookExists {
		return fmt.Errorf("book with ID %d does not exist", bookID)
	}
	if member, memberExists := library.members[memberID]; memberExists {
		if book.Status == "available" {
			member.BorrowBook(book)
			book.Status = "borrowed"
			library.store_book[bookID] = book
			return nil
		}
		return fmt.Errorf("book is already borrowed")
	}
	return fmt.Errorf("member with ID %d does not exist", memberID)
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := lib.store_book[bookID]
	if !bookExists {
		return fmt.Errorf("book with ID %d does not exist", bookID)
	}
	if member, memberExists := lib.members[memberID]; memberExists {
		if book.Status == "borrowed" {
			member.ReturnBook(bookID)
			book.Status = "available" // Update the status back to available
			lib.store_book[bookID] = book
			return nil
		}
		return fmt.Errorf("book is not borrowed")
	}
	return fmt.Errorf("member with ID %d does not exist", memberID)
}

func (lib *Library) ListAvailableBooks() []models.Book {
	available_books := []models.Book{}
	for _, book := range lib.store_book {
		if book.Status == "available" {
			available_books = append(available_books, book)
		}
	}
	return available_books
}

func (lib *Library) ListBorrowedBooks(memberID int) map[int]models.Book {
	member, exists := lib.members[memberID]
	if exists {
		return member.AllBooks()
	}
	return make(map[int]models.Book)
}

func (lib *Library) getID() int {
	return len(lib.members) + 1
}

func (lib *Library) Register(memberName string) int {
	id := lib.getID()
	member := models.CreateMember(id, memberName)
	lib.members[id] = member
	return id
}

func (lib *Library) Login(memberID int) bool {
	_, exists := lib.members[memberID]
	return exists
}

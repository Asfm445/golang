package models

type Member struct {
	ID             int
	Name           string
	borrowed_books map[int]Book
}

func CreateMember(id int, name string) Member {
	newm := Member{id, name, make(map[int]Book)}
	return newm
}

func (m Member) BorrowBook(book Book) {
	m.borrowed_books[book.ID] = book
}
func (m Member) ReturnBook(bookID int) {
	delete(m.borrowed_books, bookID)
}
func (m Member) AllBooks() map[int]Book {
	return m.borrowed_books
}

package models

type BookStatus string

// Define constants for the status
const (
	Available BookStatus = "available"
	Borrowed  BookStatus = "borrowed"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Status BookStatus
}

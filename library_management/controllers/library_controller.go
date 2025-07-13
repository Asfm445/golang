package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func Control() {
	var userID int
	library := services.NewLibrary()
	state := RegisterLoginPage(&userID, library)

	for {
		if state == 2 {
			state = RegisterLoginPage(&userID, library)
		}
		if state == 0 {
			continue // If registration/login fails, restart
		} else if state == -1 {
			break // Exit the control loop
		}
		state = menu(userID, library)
	}
}

func RegisterLoginPage(userID *int, library *services.Library) int {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n--- Welcome to the Library System ---")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Print("Enter choice: ")

	if !scanner.Scan() {
		fmt.Println("Failed to read input")
		return 0
	}
	choiceStr := scanner.Text()
	choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
	if err != nil {
		fmt.Println("Invalid input")
		return 0
	}

	switch choice {
	case 1:
		fmt.Print("Enter your ID: ")
		if !scanner.Scan() {
			fmt.Println("Failed to read input")
			return 0
		}
		id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Invalid user ID")
			return 0
		}
		*userID = id
		if !library.Login(*userID) {
			fmt.Println("You have not registered to this library.")
			return 0
		}
	case 2:
		fmt.Print("Enter your name: ")
		if !scanner.Scan() {
			fmt.Println("Failed to read name")
			return 0
		}
		name := strings.TrimSpace(scanner.Text())
		*userID = library.Register(name)
		fmt.Println("your id: ", *userID)
		return 1
	default:
		fmt.Println("Invalid choice")
		return 0
	}

	return 1
}

func menu(userID int, library *services.Library) int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Enter new book")
		fmt.Println("2. Remove book")
		fmt.Println("3. Borrow book")
		fmt.Println("4. Return book")
		fmt.Println("5. List available books")
		fmt.Println("6. List borrowed books")
		fmt.Println("7. Logout")
		fmt.Println("8. Exit")
		fmt.Print("Enter your choice: ")

		if !scanner.Scan() {
			fmt.Println("Failed to read input")
			continue
		}
		choiceStr := scanner.Text()
		choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			var book models.Book
			fmt.Print("Enter Book Title: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read title")
				continue
			}
			book.Title = strings.TrimSpace(scanner.Text())

			fmt.Print("Enter Book Author: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read author")
				continue
			}
			book.Author = strings.TrimSpace(scanner.Text())

			library.AddBook(book)
			fmt.Println("Book added successfully.")

		case 2:
			fmt.Print("Enter Book ID to remove: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				continue
			}
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid book ID")
				continue
			}
			library.RemoveBook(id)
			fmt.Println("Book removed if it existed.")

		case 3:
			fmt.Print("Enter Book ID to borrow: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				continue
			}
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid book ID")
				continue
			}
			library.BorrowBook(id, userID)

		case 4:
			fmt.Print("Enter Book ID to return: ")
			if !scanner.Scan() {
				fmt.Println("Failed to read input")
				continue
			}
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid book ID")
				continue
			}
			library.ReturnBook(id, userID)

		case 5:
			fmt.Printf("\n%-5s %-20s %-20s\n", "ID", "Book Title", "Book Author")
			for _, book := range library.ListAvailableBooks() {
				fmt.Printf("%-5d %-20s %-20s\n", book.ID, book.Title, book.Author)
			}

		case 6:
			fmt.Printf("\n%-5s %-20s %-20s\n", "ID", "Book Title", "Book Author")
			for _, book := range library.ListBorrowedBooks(userID) {
				fmt.Printf("%-5d %-20s %-20s\n", book.ID, book.Title, book.Author)
			}

		case 7:
			fmt.Println("Logging out...")
			return 2

		case 8:
			fmt.Println("Exiting program...")
			return -1

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
